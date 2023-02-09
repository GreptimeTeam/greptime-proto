# greptime-proto

GreptimeDB protobuf files.

## A Tour of GreptimeDB's Protobuf Files

You can find all protobuf files in the directory "proto" under the project's root.

Here's a brief introduction to some protobuf files:

```text
~ tree
.
├── greptime
│   └── v1
│       ├── column.proto        // Definition of the "column" schema and value of a table data, returned by GreptimeDB's gRPC service.
│       ├── database.proto      // Definition of the "request" to be carried in Arrow Flight's "ticket" (yes, we are using the "[Arrow Flight RPC](https://arrow.apache.org/docs/format/Flight.html)").
│       ├── ddl.proto           // Companion to the "request" above, defines the "Data Definition Language" requests. 
│       └── meta                
│           ├── cluster.proto   // Meta cluster gRPC service.
│           ├── common.proto    // Common data structures Meta used in its gRPC services.
│           ├── heartbeat.proto // Heartbeat gRPC service. Heartbeats are submitted by Datanodes.
│           ├── route.proto     // Table route gRPC service, where distributed tables manage their partition rules.
│           └── store.proto     // Remote KV storage gRPC service.
└── prometheus
    └── remote                  // Definition for Prometheus remote read/write protocol.
        ├── remote.proto
        └── types.proto
```

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