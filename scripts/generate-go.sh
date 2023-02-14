#! /usr/bin/env bash

set -e

PROTOC=./bin/protoc/bin/protoc
GO_OUTPUT=./go

if ! [ -f "${PROTOC}" ]; then
    echo "${PROTOC} not found! Please run scripts/install-protoc.sh first."
    exit 1
fi

# TODO(zyy17): Can we make the following commands as one command?
${PROTOC} --go_out=${GO_OUTPUT} --go_opt=paths=source_relative \
          --go-grpc_out=${GO_OUTPUT}  --go-grpc_opt=paths=source_relative \
          -Iproto proto/greptime/v1/meta/*.proto

${PROTOC} --go_out=${GO_OUTPUT}  --go_opt=paths=source_relative \
          --go-grpc_out=${GO_OUTPUT}  --go-grpc_opt=paths=source_relative \
          -Iproto proto/greptime/v1/*.proto

${PROTOC} --go_out=${GO_OUTPUT}  --go_opt=paths=source_relative \
          --go-grpc_out=${GO_OUTPUT}  --go-grpc_opt=paths=source_relative \
          -Iproto proto/prometheus/remote/*.proto
