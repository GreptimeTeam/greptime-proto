#! /usr/bin/env bash

set -e

PROTO_ROOT=./proto
GO_OUTPUT=./go

protoc -I=${PROTO_ROOT} -I /opt/include --go_out=paths=source_relative:${GO_OUTPUT} $(find  ./proto/  -type f  -iname "*.proto")
protoc -I=${PROTO_ROOT} -I /opt/include \
       --go-grpc_out=paths=source_relative:${GO_OUTPUT} \
       $(find ./proto/ -type f -iname "*.proto")

