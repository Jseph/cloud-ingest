package jcp

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"time"
)

var (
	prodServerAddr = "josephcox-storagetransfer.sandbox.googleapis.com:443"
)

func NewConnection() (*grpc.ClientConn, error) {
	ctx := context.Background()
	oauthCreds, err := oauth.NewApplicationDefault(ctx)
	if err != nil {
		return nil, err
	}
	return grpc.Dial(prodServerAddr,
		grpc.WithBlock(),
		grpc.WithTimeout(10*time.Second),
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		grpc.WithPerRPCCredentials(oauthCreds),
	)
}

func NewFakeConnection() (*grpc.ClientConn, error) {
	server := NewFakeServer()
	return grpc.Dial(
		server.Addr,
		grpc.WithInsecure(),
	)
}

func NewServerConnection(server *Server) (*grpc.ClientConn, error) {
	return grpc.Dial(
		server.Addr,
		grpc.withInsecure(),
	)
}
