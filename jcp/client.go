package jcp

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"github.com/golang/protobuf/proto"
	transferpb "github.com/GoogleCloudPlatform/cloud-ingest/proto/transfer_go_proto"
)

var (
	timeout = 10 * time.Second
)

type Client struct {
	client transferpb.StorageTransferServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		client: transferpb.NewStorageTransferServiceClient(conn),
	}
}

func (c *Client) GetGoogleServiceAccount(projectId string) (string, error){
	request := &transferpb.GetGoogleServiceAccountRequest{
		ProjectId: projectId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.client.GetGoogleServiceAccount(ctx, request)
	if err != nil {
		return "", err
	}
	return resp.AccountEmail, nil
}

func (c *Client) ReportTaskProgress(r *transferpb.ReportTaskProgressRequest) (*transferpb.ReportTaskProgressResponse, error) {
	fmt.Printf("Reporting task progress %v\n\n%v", proto.MarshalTextString(r), c)
	fmt.Printf("%v, %v", c, c == nil)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.client.ReportTaskProgress(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
