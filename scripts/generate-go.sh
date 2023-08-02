#! /usr/bin/env bash

set -e

GO_OUTPUT=./go

# TODO(zyy17): Can we make the following commands as one command?
protoc --go_out=${GO_OUTPUT} --go_opt=paths=source_relative \
       --go-grpc_out=${GO_OUTPUT}  --go-grpc_opt=paths=source_relative \
       -Iproto proto/greptime/v1/meta/*.proto

protoc --go_out=${GO_OUTPUT}  --go_opt=paths=source_relative \
       --go-grpc_out=${GO_OUTPUT}  --go-grpc_opt=paths=source_relative \
       -Iproto proto/greptime/v1/*.proto

protoc --go_out=${GO_OUTPUT}  --go_opt=paths=source_relative \
       --go-grpc_out=${GO_OUTPUT}  --go-grpc_opt=paths=source_relative \
       -Iproto proto/prometheus/remote/*.proto

protoc --go_out=${GO_OUTPUT}  --go_opt=paths=source_relative \
       --go-grpc_out=${GO_OUTPUT}  --go-grpc_opt=paths=source_relative \
       -Iproto proto/substrait_extension/*.proto

protoc --go_out=${GO_OUTPUT} --go_opt=paths=source_relative \
       --go-grpc_out=${GO_OUTPUT}  --go-grpc_opt=paths=source_relative \
       -Iproto proto/greptime/v1/mito/*.proto
