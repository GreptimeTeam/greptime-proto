#! /usr/bin/env bash

PROTO_ROOT=./proto
JAVA_OUTPUT=./java/src/main/java/

protoc -I=${PROTO_ROOT} -I /opt/include --java_out=${JAVA_OUTPUT} $(find  ./proto/  -type f  -iname "*.proto")
