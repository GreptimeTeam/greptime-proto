#!/bin/bash

PROTO_ROOT=../../../../proto

protoc -I=${PROTO_ROOT} --java_out=../java/ \
  ${PROTO_ROOT}/greptime/v1/column.proto \
  ${PROTO_ROOT}/greptime/v1/ddl.proto \
  ${PROTO_ROOT}/greptime/v1/database.proto
