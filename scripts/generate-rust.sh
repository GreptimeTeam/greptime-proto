#! /usr/bin/env bash

set -e

export PROTOC_INCLUDE=/opt/include

cargo run --manifest-path xtask/Cargo.toml -- generate-rust
