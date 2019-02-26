package agent

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"

	taskpb "github.com/GoogleCloudPlatform/cloud-ingest/proto/task_go_proto"
	transferpb "github.com/GoogleCloudPlatform/cloud-ingest/proto/transfer_go_proto"
	codepb "google.golang.org/genproto/googleapis/rpc/code"
)

var (
	versionKey = "serialized_version"
)

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Base64 encoding avoids having to consider non-utf-8 valid bytes in the serialized proto
// or potentially splitting up a valid multi-character utf-8 rune. This results in a 1.3x blowup
// in serialized task size.  If this is a problem we can consider changing TaskSpec to hold []byte
// rather than string.
func Pack(taskReqMsg *taskpb.TaskReqMsg, taskSpec *transferpb.TaskSpec) error {
	taskSpec.TaskProperties = []*transferpb.TaskProperty{
		&transferpb.TaskProperty{
			Name:  versionKey,
			Value: "0",
		},
	}
	data, err := proto.Marshal(taskReqMsg)
	if err != nil {
		return err
	}
	encoding := base64.StdEncoding
	var encoded string
	encoded = encoding.EncodeToString(data)
	for i := 0; i*1024 < len(encoded); i++ {
		taskSpec.TaskProperties = append(taskSpec.TaskProperties, &transferpb.TaskProperty{
			Name:  strconv.Itoa(i),
			Value: encoded[i*1024 : min((i+1)*1024, len(encoded))],
		})
	}
	return nil
}

func Unpack(taskSpec *transferpb.TaskSpec, taskReqMsg *taskpb.TaskReqMsg) error {
	m := make(map[string]string)
	for _, taskProperty := range taskSpec.TaskProperties {
		m[taskProperty.Name] = taskProperty.Value
	}
	fmt.Println("taskPropertyMap:", m)
	if m[versionKey] != "0" {
		errorString := fmt.Sprintf("error unpacking TaskSpec, expected %s = 0 but found %s",
			versionKey, m[versionKey])
		glog.Warning(errorString)
		return errors.New(errorString)
	}
	var sb strings.Builder
	for i := 0; i < len(m)-1; i++ {
		sb.WriteString(m[strconv.Itoa(i)])
	}
	encoding := base64.StdEncoding
	decoded, err := encoding.DecodeString(sb.String())
	if err != nil {
		return err
	}
	return proto.Unmarshal(decoded, taskReqMsg)
}

func createLogEntries(taskSpec *transferpb.TaskSpec, log *taskpb.Log) []transferpb.LogEntry {
	return []transferpb.LogEntry{
		transferpb.LogEntry{
			ProjectId:             taskSpec.ProjectId,
			TransferJobName:       taskSpec.Name,
			TransferOperationName: taskSpec.TransferOperationName,
			TaskName:              taskSpec.Name,
		},
	}
}

func newProcessListTask(jobRun string, respMsg *taskpb.TaskRespMsg) *transferpb.TaskSpec {
	listSpec := respMsg.ReqSpec.GetListSpec()
	processListSpec := taskpb.ProcessListSpec{
		DstListResultBucket: listSpec.DstListResultBucket,
		DstListResultObject: listSpec.DstListResultObject,
		SrcDirectory: listSpec.SrcDirectories[0],
	}
	reqMsg := taskpb.TaskReqMsg{
		TaskRelRsrcName: respMsg.TaskRelRsrcName,
		JobrunRelRsrcName: jobRun,
		JobRunVersion: respMsg.JobRunVersion,
		Spec: &taskpb.Spec{
			Spec: &taskpb.Spec_ProcessListSpec{
				ProcessListSpec: &processListSpec,
			},
		},
	}
	var taskSpec transferpb.TaskSpec
	Pack(&reqMsg, &taskSpec)
	taskSpec.TaskType = "PROCESS_LIST"
	return &taskSpec
}

func NewReportTaskProgressRequest(taskSpec *transferpb.TaskSpec, taskRespMsg *taskpb.TaskRespMsg) *transferpb.ReportTaskProgressRequest {
	req := &transferpb.ReportTaskProgressRequest{
		Name:                  taskSpec.Name,
		ProjectId:             taskSpec.ProjectId,
		TransferOperationName: taskSpec.TransferOperationName,
		WorkerId:              "worker_0",
		LeaseTokenId:          taskSpec.LeaseTokenId,
	}
	if taskRespMsg.Status == "FAILURE" {
		req.TaskStatus = transferpb.TaskStatus_NONRETRIABLE_FAILURE
		var errorCode codepb.Code
		switch taskRespMsg.FailureType {
		case taskpb.FailureType_FILE_MODIFIED_FAILURE,
			taskpb.FailureType_HASH_MISMATCH_FAILURE,
			taskpb.FailureType_PRECONDITION_FAILURE:
			errorCode = codepb.Code_FAILED_PRECONDITION
		case taskpb.FailureType_FILE_NOT_FOUND_FAILURE:
			errorCode = codepb.Code_NOT_FOUND
		case taskpb.FailureType_PERMISSION_FAILURE:
			errorCode = codepb.Code_PERMISSION_DENIED
		default:
			errorCode = codepb.Code_UNKNOWN
		}
		req.ErrorSummaries = []*transferpb.ErrorSummary{
			&transferpb.ErrorSummary{
				ErrorCode:  errorCode,
				ErrorCount: 1,
				ErrorLogEntries: []*transferpb.ErrorLogEntry{
					&transferpb.ErrorLogEntry{
						Url: "http://a.a",
						ErrorDetails: []string{
							taskRespMsg.FailureMessage},
					},
				},
			}}
	} else {
		//TODO: add counter support here.
		spec := taskRespMsg.ReqSpec.Spec
		if taskRespMsg.RespSpec != nil {
			spec = taskRespMsg.RespSpec.Spec
		}
		switch spec.(type) {
		case *taskpb.Spec_ListSpec:
			req.TaskStatus = transferpb.TaskStatus_COMPLETED
			req.GeneratedTaskSpecs = []*transferpb.TaskSpec{
				newProcessListTask(taskSpec.TransferOperationName, taskRespMsg),
			}
		case *taskpb.Spec_CopySpec:
			copySpec := taskRespMsg.RespSpec.GetCopySpec()
			if copySpec.FileBytes == copySpec.BytesCopied {
				glog.Errorf("marking job COMPLETED %v = %v", copySpec.FileBytes,
					copySpec.BytesCopied)
				req.TaskStatus = transferpb.TaskStatus_COMPLETED
			} else {
				req.TaskStatus = transferpb.TaskStatus_IN_PROGRESS
				req.UpdatedTaskSpec = proto.Clone(taskSpec).(*transferpb.TaskSpec)
				Pack(&taskpb.TaskReqMsg{
					TaskRelRsrcName:   taskRespMsg.TaskRelRsrcName,
					JobrunRelRsrcName: taskSpec.TransferOperationName,
					JobRunVersion:     taskRespMsg.JobRunVersion,
					Spec:              taskRespMsg.RespSpec},
					req.UpdatedTaskSpec)
			}
		case *taskpb.Spec_CopyBundleSpec:
			req.TaskStatus = transferpb.TaskStatus_COMPLETED
		}
		req.TaskStatus = transferpb.TaskStatus_COMPLETED
	}
	// TODO: Create Log Entries
	return req
}
