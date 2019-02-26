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
	"github.com/golang/protobuf/proto"
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

type Task struct {
	Spec *transferpb.TaskSpec
	Status transferpb.TaskStatus
}

// Underlying service implementer.
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
			tasks: map[string]*Task{},
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

func (s *Server) InsertTask(spec transferpb.TaskSpec) {
	s.TServer.mu.Lock()
	defer s.TServer.mu.Unlock()
	s.TServer.tasks[spec.Name] = &Task{
		Spec: &spec,
		Status: transferpb.TaskStatus_IN_PROGRESS,
	}
}

func (s *Server) SendAllTasks(t *pubsub.Topic) error {
	s.TServer.mu.Lock()
	defer s.TServer.mu.Unlock()
	ctx := context.Background()
	for _, task := range s.TServer.tasks {
		data, err := proto.Marshal(task.Spec)
		if err != nil {
			return err
		}
		res := t.Publish(ctx, &pubsub.Message{Data: data})
		res.Get(ctx)
	}
	return nil
}

func (s *Server) GetTask(name string) *Task {
	s.TServer.mu.Lock()
	defer s.TServer.mu.Unlock()
	return s.TServer.tasks[name]
}

func (s *TServer) ReportTaskProgress(_ context.Context, r *transferpb.ReportTaskProgressRequest) (*transferpb.ReportTaskProgressResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Check request preconditions.
	task, ok := s.tasks[r.Name]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Task %v not in database", r.Name)
	}
	spec := task.Spec
	if spec.LeaseTokenId != r.LeaseTokenId {
		return nil, status.Errorf(codes.FailedPrecondition, "Lease token doesn't match")
	}
	if task.Status != transferpb.TaskStatus_IN_PROGRESS {
		return nil, status.Errorf(codes.FailedPrecondition,
		"Cannot report progress for task in state %v.",
		task.Status.String())
	}
	if spec.TransferOperationName != r.TransferOperationName {
		return nil, status.Errorf(codes.FailedPrecondition,
		"Transfer operation name doesn't match.")
	}
	// Persist changes to the database.
	if r.UpdatedTaskSpec != nil {
		*spec = *r.UpdatedTaskSpec
	}
	task.Status = r.TaskStatus
	for _, spec := range r.GeneratedTaskSpecs {
		name := strconv.Itoa(rand.Int())
		s.tasks[name] = &Task{
			Spec: proto.Clone(spec).(*transferpb.TaskSpec),
			Status: transferpb.TaskStatus_IN_PROGRESS,
		}
		s.tasks[name].Spec.LeaseTokenId = strconv.Itoa(rand.Int())
		s.tasks[name].Spec.Name = name
	}
	newLeaseId := strconv.Itoa(rand.Int())
	task.Spec.LeaseTokenId = newLeaseId
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
