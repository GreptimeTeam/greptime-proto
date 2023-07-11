#! /usr/bin/env bash

set -e

PROTO_ROOT=./proto
CPP_OUTPUT=./c++

protoc --proto_path=proto/ --cpp_out=${CPP_OUTPUT}   \
    ${PROTO_ROOT}/greptime/v1/common.proto   \
    ${PROTO_ROOT}/greptime/v1/prom.proto   \
    ${PROTO_ROOT}/greptime/v1/column.proto   \
    ${PROTO_ROOT}/greptime/v1/ddl.proto   \
    ${PROTO_ROOT}/greptime/v1/database.proto
