package agent

import (
	"testing"

	"github.com/golang/protobuf/proto"

	taskpb "github.com/GoogleCloudPlatform/cloud-ingest/proto/task_go_proto"
	transferpb "github.com/GoogleCloudPlatform/cloud-ingest/proto/transfer_go_proto"
)

func expectOk(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected no error but found %v", err)
	}
}

func expectEq(req1, req2 proto.Message, t *testing.T) {
	t.Helper()
	if !proto.Equal(req1, req2) {
		t.Errorf("Expected equality between %v and %v", req1, req2)
	}
}

func TestPackUnpackDefaultTaskReqMsg(t *testing.T) {
	var req1 taskpb.TaskReqMsg
	var req2 taskpb.TaskReqMsg
	var spec transferpb.TaskSpec
	expectOk(Pack(&req1, &spec), t)
	expectOk(Unpack(&spec, &req2), t)
	expectEq(&req1, &req2, t)
}

func TestPackUnpackPopulatdTaskReqMsg(t *testing.T) {
	req1 := taskpb.TaskReqMsg{
		TaskRelRsrcName:   "asdf",
		JobrunRelRsrcName: "qwer",
		JobRunVersion:     "zxcv",
		Spec: &taskpb.Spec{
			Spec: &taskpb.Spec_CopySpec{
				CopySpec: &taskpb.CopySpec{
					SrcFile:               "hjk",
					ExpectedGenerationNum: 34,
					DstObject:             "aiien",
					BytesCopied:           64,
				},
			},
		},
	}
	var req2 taskpb.TaskReqMsg
	var spec transferpb.TaskSpec
	expectOk(Pack(&req1, &spec), t)
	expectOk(Unpack(&spec, &req2), t)
	expectEq(&req1, &req2, t)
}

func TestNewReportTaskProgressRequestListTask(t *testing.T) {
	name := "name"
	projectId := "project_1"
	operationName := "transferOp_1"
	leaseTokenId := "some_lease"
	taskSpec := transferpb.TaskSpec{
		Name:                  name,
		ProjectId:             projectId,
		TransferOperationName: operationName,
		LeaseTokenId:          leaseTokenId,
	}
	spec := &taskpb.Spec{
		Spec: &taskpb.Spec_ListSpec{
			ListSpec: &taskpb.ListSpec{
				DstListResultBucket: "a",
				DstListResultObject: "a/b",
				SrcDirectories:      []string{"/"},
			}}}
	taskRespMsg := taskpb.TaskRespMsg{
		Status:   "SUCCESS",
		ReqSpec:  spec,
		RespSpec: spec,
	}
	expected := &transferpb.ReportTaskProgressRequest{
		Name:                  name,
		ProjectId:             projectId,
		TransferOperationName: operationName,
		WorkerId:              "worker_0",
		LeaseTokenId:          leaseTokenId,
		TaskStatus:            transferpb.TaskStatus_COMPLETED,
		GeneratedTaskSpecs: []*transferpb.TaskSpec{
			newProcessListTask(operationName, &taskRespMsg)},
	}
	res := NewReportTaskProgressRequest(&taskSpec, &taskRespMsg)
	expectEq(expected, res, t)
}
