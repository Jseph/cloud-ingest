// Provides a fake jcp server which can publish task messages to the outbound task
// queue and process ReportTaskProgress requests.
package jcp

import (
	"time"
	"context"
	"math/rand"
	"strconv"
	"net"
	"sync"
	"fmt"

	"cloud.google.com/go/pubsub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/golang/protobuf/ptypes"
	transferpb "github.com/GoogleCloudPlatform/cloud-ingest/proto/transfer_go_proto"
)

// Server is a fake jcp server.
type Server struct {
	Addr string
	Port int
	l net.Listener
	Gsrv *grpc.Server
	TServer TServer
}

// Unerlying service implementer.
type TServer struct {
	transferpb.StorageTransferServiceServer

	mu sync.Mutex
	tasks map[string]*Task
}

func NewFakeServer() *Server {
	// Attempt to find an open port.
	fmt.Printf("Creating fake JCP server\n")
	var port int
	var l net.Listener
	var err error
	for i:=0; i<5; i++ {
		port = rand.Int() % 65535
		l, err = net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
		if err == nil {
			break
		}
	}
	if err != nil {
		panic(fmt.Sprintf("jcp.NewServer: %v", err))
	}
	fmt.Printf("Found an open port:%d\n", port)
	s := &Server{
		Addr: l.Addr().String(),
		Port: port,
		l: l,
		Gsrv: grpc.NewServer(),
		TServer: TServer{
			tasks: map[string]*transferpb.TaskSpec{},
		},
	}
	transferpb.RegisterStorageTransferServiceServer(s.Gsrv, &s.TServer)
	go func () {
		fmt.Printf("Starting GRPC server\n")
		if err := s.Gsrv.Serve(s.l); err != nil {
			fmt.Errorf("jcp.NewServer: %v", err)
		}
	}()
	return s
}

func (s *Server) Close() {
	s.Gsrv.Stop()
	s.l.Close()
}

func (s *Server) InsertTask(task transferpb.TaskSpec) {
	s.TServer.mu.Lock()
	defer s.TServer.mu.Unlock()
	s.TServer.tasks[task.Name] = &task
}

func (s *Server) SendAllTasks(t *pubsub.Topic) error {
	s.TServer.mu.Lock()
	defer s.TServer.mu.Unlock()
	ctx := context.Background()
	for _, task := range s.TServer.tasks {
		data, err := proto.Marshal(task)
		if err != nil {
			return err
		}
		t.Publish(ctx, &pubsub.Message{Data: data})
		t.Get(ctx)
	}
	return nil
}

func (s *TServer) ReportTaskProgress(_ context.Context, r *transferpb.ReportTaskProgressRequest) (*transferpb.ReportTaskProgressResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Check request preconditions.
	task, ok := s.tasks[r.Name]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Task %v not in database", r.Name)
	}
	if task.LeaseTokenId != r.LeaseTokenId {
		return nil, status.Errorf(codes.FailedPrecondition, "Lease token doesn't match")
	}
	if task.Status != transferpb.TaskStatus_IN_PROGRESS {
		return nil, status.Errorf(codes.FailedPrecondition,
		"Cannot report progress for task in state %v.",
		task.Status.String())
	}
	if task.TransferOperationName != r.TransferOperationName {
		return nil, status.Errorf(codes.FailedPrecondition,
		"Transfer operation name doesn't match.")
	}
	// Persist changes to the database.
	if r.UpdatedTaskSpec != nil {
		s.tasks[r.Name] = r.UpdatedTaskSpec
	}
	for _, spec := range r.GeneratedTaskSpecs {
		name := strconv.Itoa(rand.Int())
		s.tasks[name] = proto.Clone(spec).(*transferpb.TaskSpec)
		s.tasks[name].LeaseTokenId = strconv.Itoa(rand.Int())
		s.tasks[name].Name = name
	}
	newLeaseId := strconv.Itoa(rand.Int())
	task.LeaseTokenId = newLeaseId
	// Generate the response.
	response := &transferpb.ReportTaskProgressResponse{
		LeaseTokenId: newLeaseId,
		EvictTask: false,
		NewTaskLeaseExpiration: ptypes.DurationProto(24 * time.Hour),
		ResourceGrants: r.ResourceRequests,
	}
	return response, nil
}

func (s *TServer) GetGoogleServiceAccount(_ context.Context, r *transferpb.GetGoogleServiceAccountRequest) (*transferpb.GoogleServiceAccount, error) {
  return &transferpb.GoogleServiceAccount{
	  AccountEmail: "test@google.com",
  }, nil
}
