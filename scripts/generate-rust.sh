#! /usr/bin/env bash

set -e

cargo run --manifest-path xtask/Cargo.toml -- generate-rust
