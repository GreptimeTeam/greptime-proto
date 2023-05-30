#! /usr/bin/env bash

PROTO_ROOT=./proto
JAVA_OUTPUT=./java/src/main/java/

protoc -I=${PROTO_ROOT} --java_out=${JAVA_OUTPUT} \
  ${PROTO_ROOT}/greptime/v1/column.proto \
  ${PROTO_ROOT}/greptime/v1/ddl.proto \
  ${PROTO_ROOT}/greptime/v1/database.proto
