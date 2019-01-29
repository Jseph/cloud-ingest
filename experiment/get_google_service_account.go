package main

import (
	"flag"
	"github.com/GoogleCloudPlatform/cloud-ingest/jcp"
	"google.golang.org/grpc"
	"log"
)

var (
	serverAddr = "josephcox-storagetransfer.sandbox.googleapis.com:443"
)

func main() {
	fakeServer := flag.Bool("fake", true, "If a local fake server should be used.")
	flag.Parse()

	var conn *grpc.ClientConn
	var err error
	if *fakeServer {
		conn, err = jcp.NewFakeConnection()
	} else {
		conn, err = jcp.NewConnection()
	}
	if err != nil {
		log.Fatalf("failed to establish a connection: %v", err)
	}
	defer conn.Close()
	client := jcp.NewClient(conn)
	resp, err := client.GetGoogleServiceAccount("207531450211")
	if err != nil {
		log.Fatalf("fail to get account: %v", err)
	}
	log.Println(resp)
}
