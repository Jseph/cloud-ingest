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
	"cloud.google.com/go/storage"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"path/filepath"
)

const (
	noTaskIdInListOutput string = ("expected task ID %s in first line of list task output file, but got %s")
)

type ListProgressMessageHandler struct {
	Store                Store
	ListingResultReader  ListingResultReader
	ObjectMetadataReader ObjectMetadataReader
}

func (h *ListProgressMessageHandler) retrieveGenerationNumber(bucketName string, objectName string) (int64, error) {
	metadata, err := h.ObjectMetadataReader.GetMetadata(bucketName, objectName)
	if err == nil {
		return metadata.GenerationNumber, nil
	} else if err == storage.ErrObjectNotExist {
		return 0, nil
	} else {
		return 0, err
	}
}

func (h *ListProgressMessageHandler) HandleMessage(
	jobSpec *JobSpec, taskCompletionMessage *TaskCompletionMessage) (*TaskUpdate, error) {

	taskUpdate, err := TaskCompletionMessageToTaskUpdate(taskCompletionMessage)
	if err != nil {
		log.Printf("Error extracting taskCompletionMessage %v: %v", taskCompletionMessage, err)
		return nil, err
	}

	taskUpdate.Task.TaskType = listTaskType // Set the type first.
	if taskUpdate.Task.Status != Success {
		return taskUpdate, nil
	}
	task := taskUpdate.Task

	// TODO(b/63014658): denormalize the task spec into the progress message, so
	// you do not have to query the database to get the task spec.
	// TODO(b/67420045): Message handlers should not have the store object.
	// Manipulating the store should be isolated from handling the message.
	taskSpec, err := h.Store.GetTaskSpec(task.JobConfigId, task.JobRunId, task.TaskId)
	if err != nil {
		log.Printf("Error getting task spec of task: %v, with error: %v.",
			task, err)
		return nil, err
	}

	var listTaskSpec ListTaskSpec
	if err := json.Unmarshal([]byte(taskSpec), &listTaskSpec); err != nil {
		log.Printf(
			"Error decoding task spec: %s, with error: %v.", taskSpec, err)
		return nil, err
	}

	filePaths, err := h.ListingResultReader.ReadListResult(
		listTaskSpec.DstListResultBucket, listTaskSpec.DstListResultObject)
	if err != nil {
		log.Printf(
			"Error reading the list task result, list task spec: %v, with error: %v.",
			listTaskSpec, err)
		return nil, err
	}
	taskIdFromFile := <-filePaths

	if taskIdFromFile != task.getTaskFullId() {
		return nil, errors.New(
			fmt.Sprintf(noTaskIdInListOutput, task.getTaskFullId(), taskIdFromFile))
	}
	var newTasks []*Task
	for filePath := range filePaths {
		uploadGCSTaskId := GetUploadGCSTaskId(filePath)
		dstObject, _ := filepath.Rel(listTaskSpec.SrcDirectory, filePath)

		generationNumber, err := h.retrieveGenerationNumber(jobSpec.GCSBucket, dstObject)
		if err != nil {
			log.Printf(
				"Error reading file metadata for %s:%s, err: %v.",
				jobSpec.GCSBucket, filePath, err)
			return nil, err
		}

		uploadGCSTaskSpec := UploadGCSTaskSpec{
			SrcFile:               filePath,
			DstBucket:             jobSpec.GCSBucket,
			DstObject:             dstObject,
			ExpectedGenerationNum: generationNumber,
		}
		uploadGCSTaskSpecJson, err := json.Marshal(uploadGCSTaskSpec)
		if err != nil {
			log.Printf(
				"Error encoding task spec to JSON string, task spec: %v, err: %v.",
				uploadGCSTaskSpec, err)
			return nil, err
		}
		newTasks = append(newTasks, &Task{
			JobConfigId: task.JobConfigId,
			JobRunId:    task.JobRunId,
			TaskId:      uploadGCSTaskId,
			TaskType:    uploadGCSTaskType,
			TaskSpec:    string(uploadGCSTaskSpecJson),
		})
	}
	taskUpdate.NewTasks = newTasks

	return taskUpdate, nil
}
