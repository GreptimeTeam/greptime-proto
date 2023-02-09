# greptime-proto

GreptimeDB protobuf files.

## Usage

### Rust

```Toml
# Add this repository as dependency to your Cargo.toml file:
greptime-proto = { git = "https://github.com/GreptimeTeam/greptime-proto.git" }
```

```Rust
// To use the GreptimeDB's gRPC service:
use greptime_proto::v1::*;

// To talk to GreptimeDB Meta service:
use greptime_proto::v1::meta::*;

// To request GreptimeDB as Prometheus remote read/write:
use greptime_proto::prometheus::remote::*;
```

## Requirements

### [google/protobuf](https://github.com/protocolbuffers/protobuf)

Please use protobuf v3.

## Build

```text
make
```