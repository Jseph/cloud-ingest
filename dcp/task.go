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

package dcp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/GoogleCloudPlatform/cloud-ingest/dcp/proto"
	"github.com/GoogleCloudPlatform/cloud-ingest/helpers"
)

const (
	Unqueued int64 = 0
	Queued   int64 = 1
	Failed   int64 = 2
	Success  int64 = 3

	listTaskPrefix        string = "list"
	processListTaskPrefix string = "processList"
	copyTaskPrefix        string = "copy"

	listTaskType        int64 = 1
	processListTaskType int64 = 2
	copyTaskType        int64 = 3

	idSeparator string = ":"
)

type JobSpec struct {
	OnpremSrcDirectory string `json:"onPremSrcDirectory"`
	GCSBucket          string `json:"gcsBucket"`
	GCSDirectory       string `json:"gcsDirectory"`
}

type ListTaskSpec struct {
	DstListResultBucket   string `json:"dst_list_result_bucket"`
	DstListResultObject   string `json:"dst_list_result_object"`
	SrcDirectory          string `json:"src_directory"`
	ExpectedGenerationNum int64  `json:"expected_generation_num"`
}

type ProcessListTaskSpec struct {
	DstListResultBucket string `json:"dst_list_result_bucket"`
	DstListResultObject string `json:"dst_list_result_object"`
	SrcDirectory        string `json:"src_directory"`
	ByteOffset          int64  `json:"byte_offset"`
}

type CopyTaskSpec struct {
	SrcFile               string `json:"src_file"`
	DstBucket             string `json:"dst_bucket"`
	DstObject             string `json:"dst_object"`
	ExpectedGenerationNum int64  `json:"expected_generation_num"`
}

type Task struct {
	TaskRRStruct         TaskRRStruct
	TaskSpec             string
	TaskType             int64
	Status               int64
	CreationTime         int64
	LastModificationTime int64
	FailureMessage       string
	FailureType          proto.TaskFailureType_Type
}

type TaskParams map[string]interface{}

// TaskCompletionMessage is the response type we get back from any progress queue.
type TaskCompletionMessage struct {
	TaskRRName     string                     `json:"task_id"`
	Status         string                     `json:"status"`
	FailureType    proto.TaskFailureType_Type `json:"failure_reason"`
	FailureMessage string                     `json:"failure_message"`
	LogEntry       LogEntry                   `json:"log_entry"`
	TaskParams     TaskParams                 `json:"task_params"`
}

// TaskTransactionalSemantics is an interface that captures task-specific semantics that need to be
// performed transactionally as part of a task update. This is the hook into the heart of the
// transaction.
type TaskTransactionalSemantics interface {
	// Apply applies the taskUpdate transaction semantic, it returns bool of
	// whether that task update should be processed, and error if any errors
	// occurred.
	Apply(taskUpdate *TaskUpdate) (bool, error)
}

// TaskUpdate represents a task to be updated, with it's log entry and new tasks
// generated from the update.
type TaskUpdate struct {
	// Task to be updated.
	Task *Task

	// LogEntry that is associated with task update.
	LogEntry LogEntry

	// Parameters the original task was called with.
	OriginalTaskParams TaskParams

	// NewTasks that are generated by this task update.
	NewTasks []*Task

	// TransactionalSemantics are set individually for each task. If this is nil, it's a no-op.
	TransactionalSemantics TaskTransactionalSemantics
}

// TaskUpdateCollection is a collection of task updates to be committed to the
// store.
type TaskUpdateCollection struct {
	// tasks is a map from task id to task update details.
	tasks map[TaskRRStruct]*TaskUpdate
}

// AddTaskUpdate adds a taskUpdate into the collection. If there is a task in
// the collection that has the same task id as taskUpdate but with
// different status, the statuses will be compared and only the task with the
// higher status (as defined by canChangeTaskStatus) will be added to the
// collection.
//
// For example, consider the following task updates:
//
//	taskUpdateA := &TaskUpdate{
//		Task: &Task{
//			TaskRRStruct: *NewTaskRRStruct("project", "config", "run", "list"),
//			Status:       Failed,
//		},
//	}
//
// 	taskUpdateB := &TaskUpdate{
//		Task: &Task{
//			TaskRRStruct: *NewTaskRRStruct("project", "config", "run", "list"),
//			Status:       Success,
//		},
//		NewTasks: []*Task{
//			&Task{
//				TaskRRStruct: *NewTaskRRStruct("project", "config", "run", "copy"),
//				Status:       Unqueued,
//			},
//		},
//	}
//
// Only one of taskUpdateA and taskUpdateB can exist in the collection since they
// have the same task ID.  If taskUpdateA is already in the collection and
// the caller tries to add taskUpdateB, taskUpdateA will be replaced in the
// collection by taskUpdateB because canChangeTaskStatus(taskUpdateA.Task.Status,
// taskUpdateB.Task.Status) is true (Failed -> Success).
func (tc *TaskUpdateCollection) AddTaskUpdate(taskUpdate *TaskUpdate) {
	if tc.tasks == nil {
		tc.tasks = make(map[TaskRRStruct]*TaskUpdate)
	}
	taskRRStruct := taskUpdate.Task.TaskRRStruct

	otherTaskUpdate, exists := tc.tasks[taskRRStruct]
	if !exists || canChangeTaskStatus(otherTaskUpdate.Task.Status, taskUpdate.Task.Status) {
		// This is the only task so far with this id or it is
		// more recent than any other tasks seen so far with the same id.
		tc.tasks[taskRRStruct] = taskUpdate
	}
}

func NewListTaskSpecFromMap(params TaskParams) (*ListTaskSpec, error) {
	dstBucket, ok1 := params["dst_list_result_bucket"]
	dstObject, ok2 := params["dst_list_result_object"]
	srcDir, ok3 := params["src_directory"]
	genNum, ok4 := params["expected_generation_num"]

	if ok1 && ok2 && ok3 && ok4 {
		genNum, err := helpers.ToInt64(genNum)
		if err != nil {
			return nil, err
		}
		return &ListTaskSpec{
			DstListResultBucket:   dstBucket.(string),
			DstListResultObject:   dstObject.(string),
			SrcDirectory:          srcDir.(string),
			ExpectedGenerationNum: genNum,
		}, nil
	}

	return nil, fmt.Errorf("missing params in ListTaskSpec map: %v", params)
}

func NewProcessListTaskSpecFromMap(params map[string]interface{}) (*ProcessListTaskSpec, error) {
	dstBucket, ok1 := params["dst_list_result_bucket"]
	dstObject, ok2 := params["dst_list_result_object"]
	srcDir, ok3 := params["src_directory"]
	byteOffset, ok4 := params["byte_offset"]

	if ok1 && ok2 && ok3 && ok4 {
		byteOffset, err := helpers.ToInt64(byteOffset)
		if err != nil {
			return nil, err
		}
		return &ProcessListTaskSpec{
			DstListResultBucket: dstBucket.(string),
			DstListResultObject: dstObject.(string),
			SrcDirectory:        srcDir.(string),
			ByteOffset:          byteOffset,
		}, nil
	}

	return nil, fmt.Errorf("missing params in ProcessListTaskSpec map: %v", params)
}

func NewCopyTaskSpecFromMap(params map[string]interface{}) (*CopyTaskSpec, error) {
	srcFile, ok1 := params["src_file"]
	dstBucket, ok2 := params["dst_bucket"]
	dstObject, ok3 := params["dst_object"]
	genNum, ok4 := params["expected_generation_num"]
	if ok1 && ok2 && ok3 && ok4 {
		genNum, err := helpers.ToInt64(genNum)
		if err != nil {
			return nil, err
		}
		return &CopyTaskSpec{
			SrcFile:               srcFile.(string),
			DstBucket:             dstBucket.(string),
			DstObject:             dstObject.(string),
			ExpectedGenerationNum: genNum,
		}, nil
	}

	return nil, fmt.Errorf("missing params in CopyTaskSpec map: %v", params)
}

func (tc TaskUpdateCollection) Size() int {
	return len(tc.tasks)
}

func (tc TaskUpdateCollection) GetTaskUpdate(taskRRStruct TaskRRStruct) *TaskUpdate {
	return tc.tasks[taskRRStruct]
}

func (tc TaskUpdateCollection) GetTaskUpdates() <-chan *TaskUpdate {
	c := make(chan *TaskUpdate)
	go func() {
		defer close(c)
		for _, taskUpdate := range tc.tasks {
			c <- taskUpdate
		}
	}()
	return c
}

func (tc *TaskUpdateCollection) Clear() {
	tc.tasks = make(map[TaskRRStruct]*TaskUpdate)
}

// GetProcessListTaskID returns the task id of an processList type task for
// the given bucket and object.
func GetProcessListTaskID(bucket, object string) string {
	return processListTaskPrefix + idSeparator + bucket + idSeparator + object
}

// GetCopyTaskID returns the task id of a copy type task for the given file.
func GetCopyTaskID(filePath string) string {
	// TODO(b/64038794): The task ids should be a hash of the filePath, the
	// filePath might be too long and already duplicated in the task spec.
	return copyTaskPrefix + idSeparator + filePath
}

// canChangeTaskStatus checks whether a task can be moved from a fromStatus to
// a toStatus.
func canChangeTaskStatus(fromStatus int64, toStatus int64) bool {
	// In general, tasks will flow through Unqueued -> Queued -> (Fail|Success).
	// Transitions to Success are always allowed, to gracefully handle multiple delivery cases
	// where one run succeeded. We rely on task-specific semantics to guarantee that a Success is
	// robust (hasn't been clobbered by other activity behind the scenes).
	// Transitions to Unqueued are allowed, for reissuing tasks as needed.
	if fromStatus == Success {
		// Success is immutable.
		return false
	}

	return toStatus == Unqueued || toStatus > fromStatus
}

// constructPubSubTaskMsg constructs the Pub/Sub message for the passed task to
// send to the worker agents.
func constructPubSubTaskMsg(task *Task) ([]byte, error) {
	taskParams := make(TaskParams)
	if err := json.Unmarshal([]byte(task.TaskSpec), &taskParams); err != nil {
		return nil, fmt.Errorf(
			"error decoding JSON task spec string %s for task %v.",
			task.TaskSpec, task)
	}

	taskMsg := make(map[string]interface{})
	taskMsg["task_id"] = task.TaskRRStruct.String()
	taskMsg["task_params"] = taskParams
	return json.Marshal(taskMsg)
}

func TaskCompletionMessageFromJson(msg []byte) (*TaskCompletionMessage, error) {
	var taskCompletionMessage TaskCompletionMessage
	decoder := json.NewDecoder(bytes.NewReader(msg))
	decoder.UseNumber()
	err := decoder.Decode(&taskCompletionMessage)
	if err != nil {
		return nil, err
	}
	return &taskCompletionMessage, nil
}

// TaskCompletionMessageToTaskUpdate is a utility to generate a base TaskUpdate from a
// TaskCompletionMessage.
func TaskCompletionMessageToTaskUpdate(taskCompletionMessage *TaskCompletionMessage) (*TaskUpdate, error) {
	if taskCompletionMessage == nil {
		return nil, errors.New("received nil taskCompletionMessage")
	}

	// Construct the task.
	taskRRStruct, err := TaskRRStructFromTaskRRName(taskCompletionMessage.TaskRRName)
	if err != nil {
		return nil, err
	}
	task := &Task{
		TaskRRStruct: *taskRRStruct,
	}
	if taskCompletionMessage.Status == "SUCCESS" {
		task.Status = Success
	} else {
		task.Status = Failed
		task.FailureType = taskCompletionMessage.FailureType
		task.FailureMessage = taskCompletionMessage.FailureMessage
	}

	return &TaskUpdate{
		Task:               task,
		LogEntry:           taskCompletionMessage.LogEntry,
		OriginalTaskParams: taskCompletionMessage.TaskParams,
	}, nil
}
