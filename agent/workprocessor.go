/*
Copyright 2017 Google Inc. All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package agent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/GoogleCloudPlatform/cloud-ingest/agent/rate"
	"github.com/GoogleCloudPlatform/cloud-ingest/agent/stats"
	"github.com/GoogleCloudPlatform/cloud-ingest/jcp"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"

	taskpb "github.com/GoogleCloudPlatform/cloud-ingest/proto/task_go_proto"
	transferpb "github.com/GoogleCloudPlatform/cloud-ingest/proto/transfer_go_proto"
)

// WorkHandler is an interface to handle different task types.
type WorkHandler interface {
	// Do handles the TaskReqMsg and returns a TaskRespMsg.
	Do(ctx context.Context, taskReqMsg *taskpb.TaskReqMsg) *taskpb.TaskRespMsg
}

// WorkProcessor processes tasks of a certain type. It listens to subscription
// WorkSub, delegates to the Handler to do the work, and send progress messages
// to ProgressTopic.
type WorkProcessor struct {
	WorkSub       *pubsub.Subscription
	ProgressTopic *pubsub.Topic
	Handlers      *HandlerRegistry
	JcpClient     *jcp.Client
	StatsTracker  *stats.Tracker
}

func (wp *WorkProcessor) processMessage(ctx context.Context, msg *pubsub.Message) {
	// Data will be 1 of 2 forms, either a taskpb.TaskReqMsg or a transferpb.TaskSpec.
	// transferpb.TaskSpec messages can be converted to taskpb.TaskReqMsgs.
	var taskReqMsg taskpb.TaskReqMsg
	var taskSpec transferpb.TaskSpec
	useJcp := true
	err := proto.Unmarshal(msg.Data, &taskSpec)
	if err == nil {
		fmt.Printf("Got taskSpec %v", proto.MarshalTextString(&taskSpec))
		err = Unpack(&taskSpec, &taskReqMsg)
	}
	if err != nil {
		glog.Info("Could not parse message as jcp message.  Using old message format")
		if err = proto.Unmarshal(msg.Data, &taskReqMsg); err != nil {
			glog.Errorf("could not decode pubsub message.")
			msg.Ack()
			return
		}
		useJcp = false
	}
	fmt.Printf("Got taskReqMsg %v", proto.MarshalTextString(&taskReqMsg))
	var jcpWg sync.WaitGroup
	quitHeartbeat := make(chan bool)
	if useJcp {
		msg.Ack()
		// Perform Checkpoint 0.
		request := transferpb.ReportTaskProgressRequest{
			Name:                  taskSpec.Name,
			ProjectId:             taskSpec.ProjectId,
			TransferOperationName: taskSpec.TransferOperationName,
			WorkerId:              "worker_0",
			LeaseTokenId:          taskSpec.LeaseTokenId,
			TaskStatus:            transferpb.TaskStatus_IN_PROGRESS,
		}

		resp, err := wp.JcpClient.ReportTaskProgress(&request)
		if err != nil {
			glog.Errorf("error with checkpoint 0 %v", err)
			return
		}
		if resp.EvictTask {
			glog.Infof("recieved task eviction for %s", taskSpec.Name)
			return
		}
		taskSpec.LeaseTokenId = resp.LeaseTokenId
		// Set up a periodic heartbeat to not lose the lease.
		jcpWg.Add(1)
		go func() {
			for {
				select {
				case <-quitHeartbeat:
					jcpWg.Done()
					return
				case <-time.After(5 * time.Second):
					request := transferpb.ReportTaskProgressRequest{
						Name:                  taskSpec.Name,
						ProjectId:             taskSpec.ProjectId,
						TransferOperationName: taskSpec.TransferOperationName,
						WorkerId:              "worker_0",
						LeaseTokenId:          taskSpec.LeaseTokenId,
						TaskStatus:            transferpb.TaskStatus_IN_PROGRESS,
					}

					resp, err := wp.JcpClient.ReportTaskProgress(&request)
					if err != nil {
						glog.Errorf("error with heartbeat %v", err)
						return
					}
					taskSpec.LeaseTokenId = resp.LeaseTokenId
				}
			}
		}()
	}
	var taskRespMsg *taskpb.TaskRespMsg
	if useJcp || rate.IsJobRunActive(taskReqMsg.JobrunRelRsrcName) {
		handler, agentErr := wp.Handlers.HandlerForTaskReqMsg(&taskReqMsg)
		if agentErr != nil {
			taskRespMsg = buildTaskRespMsg(&taskReqMsg, nil, nil, *agentErr)
		} else {
			start := time.Now()
			taskRespMsg = handler.Do(ctx, &taskReqMsg)
			if wp.StatsTracker != nil {
				wp.StatsTracker.RecordTaskRespDuration(taskRespMsg, time.Now().Sub(start))
			}
		}
	} else {
		taskRespMsg = buildTaskRespMsg(&taskReqMsg, nil, nil, AgentError{
			Msg:         fmt.Sprintf("job run %s is not active", taskReqMsg.JobrunRelRsrcName),
			FailureType: taskpb.FailureType_NOT_ACTIVE_JOBRUN,
		})
	}

	if !proto.Equal(taskReqMsg.Spec, taskRespMsg.ReqSpec) {
		glog.Errorf("taskRespMsg.ReqSpec = %v, want: %v", taskRespMsg.ReqSpec, taskReqMsg.Spec)
		// The taskRespMsg.ReqSpec must equal the taskReqMsg.Spec. This is an Agent
		// coding error, do not ack the PubSub message.
		return
	}
	if ctx.Err() == context.Canceled {
		glog.Errorf("Context is canceled, not sending taskRespMsg: %v", taskRespMsg)
		// When the context is canceled midway through processing a request, instead of
		// surfacing an error which propagates to the DCP just don't send the response.
		// The work will remain on PubSub and eventually be taken up by another worker.
		return
	}
	if !useJcp {
		serializedTaskRespMsg, err := proto.Marshal(taskRespMsg)
		if err != nil {
			glog.Errorf("Cannot marshal pb %+v with err %v", taskRespMsg, err)
			// This may be a transient error, will not Ack the messages to retry again
			// when the message redelivered.
			return
		}
		pubResult := wp.ProgressTopic.Publish(ctx, &pubsub.Message{Data: serializedTaskRespMsg})
		if _, err := pubResult.Get(ctx); err != nil {
			glog.Errorf("Can not publish progress message with err: %v", err)
			// Don't ack the messages, retry again when the message is redelivered.
			return
		}
		msg.Ack()
	} else {
		// Wait for the heartbeat to finish, we don't want to have racing update requests.
		quitHeartbeat <- true
		jcpWg.Wait()

		// Convert the taskRespMsg to a ReportTaskProgressRequest.
		request := NewReportTaskProgressRequest(&taskSpec, taskRespMsg)
		resp, err := wp.JcpClient.ReportTaskProgress(request)
		if err != nil {
			glog.Errorf("Error reporting task progress: %v", err)
			return
		}
		if request.UpdatedTaskSpec != nil {
			newSpec := request.UpdatedTaskSpec
			newSpec.LeaseTokenId = resp.LeaseTokenId
			var newMsg pubsub.Message
			data, err := proto.Marshal(newSpec)
			if err != nil {
				glog.Errorf("Error marshaling new TaskSpec: %v", err)
				return
			}
			newMsg.Data = data
			wp.processMessage(ctx, &newMsg)
		}
	}

}

func (wp *WorkProcessor) Process(ctx context.Context) {
	// Use the DefaultReceiveSettings, which is ReceiveSettings{
	// 	 MaxExtension:           10 * time.Minute,
	// 	 MaxOutstandingMessages: 1000,
	// 	 MaxOutstandingBytes:    1e9,
	// 	 NumGoroutines:          1,
	// }
	// The default settings should be safe, because of the following reasons:
	// * MaxExtension: DCP should not publish messages that estimated to take more
	//                 than 10 mins.
	// * MaxOutstandingMessages: It's also capped by the memory, and this will speed
	//                           up processing of small files.
	// * MaxOutstandingBytes: 1GB memory should not be a problem for a modern machine.
	// * NumGoroutines: Does not need more than 1 routine to pull Pub/Sub messages.
	//
	// Note that the main function can override those values when constructing the
	// WorkProcessor instance.
	err := wp.WorkSub.Receive(ctx, wp.processMessage)

	if ctx.Err() != nil {
		glog.Warningf(
			"Error receiving work messages for subscription %v, with context error: %v.",
			wp.WorkSub, ctx.Err())
	}

	// The Pub/Sub client libraries already retries on retriable errors. Panic
	// here on non-retriable errors.
	if err != nil {
		glog.Fatalf("Error receiving work messages for subscription %v, with error: %v.",
			wp.WorkSub, err)
	}
}
