#!/bin/sh
# This script generates the Go protocol buffers for the protos in this dir.
protoc --go_out=task_go_proto/ task.proto
protoc --go_out=pulse_go_proto/ pulse.proto
protoc --go_out=control_go_proto/ control.proto
protoc --go_out=listfile_go_proto/ listfile.proto
protoc -I . -I ~/go/src/ --go_out=plugins=grpc:transfer_go_proto transfer.proto
