#! /usr/bin/env bash

set -e

PROTOC_GEN_GO_VERSION=v1.28.1
PROTOC_GEN_GO_GRPC_VERSION=v1.2.0

# Install protoc-gen-go and protoc-gen-go-grpc.
# Reference: https://grpc.io/docs/languages/go/quickstart/.
go install google.golang.org/protobuf/cmd/protoc-gen-go@${PROTOC_GEN_GO_VERSION}
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@${PROTOC_GEN_GO_GRPC_VERSION}
