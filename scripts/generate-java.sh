#! /usr/bin/env bash

PROTO_ROOT=./proto
JAVA_OUTPUT=./java/src/main/java/

protoc -I=${PROTO_ROOT} --java_out=${JAVA_OUTPUT} $(find  ./proto/  -type f  -iname "*.proto")
