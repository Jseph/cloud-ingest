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
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/cloud-ingest/helpers"
	"github.com/golang/mock/gomock"
)

func processListCompletionMessage() *TaskCompletionMessage {
	return &TaskCompletionMessage{
		TaskRRName: testTaskRRName,
		Status:     "SUCCESS",
		TaskParams: map[string]interface{}{
			"dst_list_result_bucket": "bucket1",
			"dst_list_result_object": "object",
			"src_directory":          "dir",
			"byte_offset":            0,
		},
		LogEntry: map[string]interface{}{"logkey": "logval"},
	}
}

func TestProcessListMessageHandlerInvalidCompletionMessage(t *testing.T) {
	handler := ProcessListMessageHandler{}

	taskCompletionMessage := processListCompletionMessage()
	taskCompletionMessage.TaskRRName = "garbage"
	_, err := handler.HandleMessage(nil, taskCompletionMessage)
	if err == nil {
		t.Error("error is nil, expected error: can not parse full task id...")
	} else if !strings.Contains(err.Error(), "cannot parse") {
		t.Errorf("expected error: %s, found: %s.", "can not parse full task id",
			err.Error())
	}
}

func TestProcessListMessageHandlerFailReadingListResult(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockListReader := NewMockListingResultReader(mockCtrl)
	errorMsg := "Failed reading listing result."
	mockListReader.EXPECT().ReadEntries(
		context.Background(), "bucket1", "object", int64(0), maxEntriesToProcess).
		Return(nil, int64(0), errors.New(errorMsg))

	handler := ProcessListMessageHandler{
		ListingResultReader: mockListReader,
	}

	_, err := handler.HandleMessage(nil, processListCompletionMessage())
	if err == nil {
		t.Errorf("error is nil, expected error: %s.", errorMsg)
	} else if err.Error() != errorMsg {
		t.Errorf("expected error: %s, found: %s.", errorMsg, err.Error())
	}
}

func TestProcessListMessageHandlerSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Setup the ListingResultReader.
	mockListReader := NewMockListingResultReader(mockCtrl)
	var newBytesProcessed int64 = 123
	entries := []ListFileEntry{
		ListFileEntry{true, "dir/dir0"},
		ListFileEntry{false, "dir/file0"},
		ListFileEntry{false, "dir/file1"},
	}
	mockListReader.EXPECT().ReadEntries(
		context.Background(), "bucket1", "object", int64(0), maxEntriesToProcess).
		Return(entries, newBytesProcessed, io.EOF)

	taskRRStruct := NewTaskRRStruct(testProjectID, testJobConfigID, testJobRunID, testTaskID)
	processListTask := &Task{
		TaskRRStruct: *taskRRStruct,
		TaskType:     processListTaskType,
		TaskSpec: `{
			"dst_list_result_bucket": "bucket1",
			"dst_list_result_object": "object",
			"src_directory": "dir",
			"byte_offset": 0
		}`,
		Status: Success,
	}

	handler := ProcessListMessageHandler{
		ListingResultReader: mockListReader,
	}

	jobSpec := &JobSpec{
		OnpremSrcDirectory: "dir",
		GCSBucket:          "bucket2",
	}

	taskUpdate, err := handler.HandleMessage(jobSpec, processListCompletionMessage())
	if err != nil {
		t.Errorf("expecting success, found error: %v.", err)
	}

	expectedTaskUpdate := &TaskUpdate{
		Task: processListTask,
		LogEntry: LogEntry{
			"endingOffset":     newBytesProcessed,
			"entriesProcessed": int64(3),
			"startingOffset":   int64(0),
		},
		OriginalTaskParams: TaskParams{
			"byte_offset":            0,
			"dst_list_result_bucket": "bucket1",
			"dst_list_result_object": "object",
			"src_directory":          "dir",
		},
		TransactionalSemantics: ProcessListingFileSemantics{
			ExpectedByteOffset:         int64(0),
			ByteOffsetForNextIteration: newBytesProcessed,
		},
	}

	// No task spec on TaskUpdate.
	processListTask.TaskSpec = ""

	if len(taskUpdate.NewTasks) != 3 {
		t.Errorf("expecting 3 tasks, found %d.", len(taskUpdate.NewTasks))
	}

	// Validate the new list task.
	{
		// Handle the task spec JSON separately, since it doesn't play well with equality checks.
		expectedNewTaskSpec := fmt.Sprintf(`{
				"dst_list_result_bucket": "bucket2",
				"dst_list_result_object": "%s/listfiles/%s/%s/dir0/list",
				"src_directory": "dir/dir0",
				"expected_generation_num": 0
			}`, cloudIngestWorkingSpace, testJobConfigID, testJobRunID)
		if !helpers.AreEqualJSON(expectedNewTaskSpec, taskUpdate.NewTasks[0].TaskSpec) {
			t.Errorf("expected task spec: %s, found: %s", expectedNewTaskSpec, taskUpdate.NewTasks[0].TaskSpec)
		}
		taskUpdate.NewTasks[0].TaskSpec = "" // Blow it away.
		expectedNewTask := &Task{            // Add task (sans spec) to our expected update.
			TaskRRStruct: TaskRRStruct{
				JobRunRRStruct: taskRRStruct.JobRunRRStruct,
				TaskID:         GetListTaskID("dir/dir0"),
			},
			TaskType: listTaskType,
		}
		expectedTaskUpdate.NewTasks = append(expectedTaskUpdate.NewTasks, expectedNewTask)
	}

	// Validate the new copy tasks.
	for i := 0; i < 2; i++ {
		// Handle the task spec JSON separately, since it doesn't play well with equality checks.
		expectedNewTaskSpec := fmt.Sprintf(`{
				"dst_bucket": "bucket2",
				"dst_object": "file%d",
				"expected_generation_num": 0,
				"src_file": "dir/file%d"
			}`, i, i)

		if !helpers.AreEqualJSON(expectedNewTaskSpec, taskUpdate.NewTasks[i+1].TaskSpec) {
			t.Errorf("expected task spec: %s, found: %s", expectedNewTaskSpec, taskUpdate.NewTasks[i+1].TaskSpec)
		}
		taskUpdate.NewTasks[i+1].TaskSpec = "" // Blow it away.
		expectedNewTask := &Task{              // Add task (sans spec) to our expected update.
			TaskRRStruct: TaskRRStruct{
				JobRunRRStruct: taskRRStruct.JobRunRRStruct,
				TaskID:         GetCopyTaskID("dir/file" + strconv.Itoa(i)),
			},
			TaskType: copyTaskType,
		}
		expectedTaskUpdate.NewTasks = append(expectedTaskUpdate.NewTasks, expectedNewTask)
	}

	DeepEqualCompare("TaskUpdate", expectedTaskUpdate, taskUpdate, t)

	// Check pieces one at a time, for convenient visualization.
	DeepEqualCompare("TaskUpdate.Task", expectedTaskUpdate.Task, taskUpdate.Task, t)
	DeepEqualCompare("TaskUpdate.LogEntry", expectedTaskUpdate.LogEntry, taskUpdate.LogEntry, t)
	DeepEqualCompare("TaskUpdate.OriginalTaskParams",
		expectedTaskUpdate.OriginalTaskParams, taskUpdate.OriginalTaskParams, t)
	for i := 0; i < 3; i++ {
		DeepEqualCompare("NewTasks", expectedTaskUpdate.NewTasks[i], taskUpdate.NewTasks[i], t)
	}
}