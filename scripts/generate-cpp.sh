#! /usr/bin/env bash

set -e

PROTO_ROOT=./proto
CPP_OUTPUT=./c++

protoc --proto_path=proto/ --cpp_out=${CPP_OUTPUT} $(find  ./proto/  -type f  -iname "*.proto")