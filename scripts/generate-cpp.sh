#! /usr/bin/env bash

set -e

PROTO_ROOT=./proto
CPP_OUTPUT=./c++

protoc -I ${PROTO_ROOT} -I /opt/include --cpp_out=${CPP_OUTPUT} $(find  ./proto/  -type f  -iname "*.proto")
protoc -I ${PROTO_ROOT} -I /opt/include --grpc_out=${CPP_OUTPUT} --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` $(find  ./proto/  -type f  -iname "*.proto")
