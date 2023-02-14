# greptime-proto

GreptimeDB protobuf files.

## Build

### Requirement

- [google/protobuf](https://github.com/protocolbuffers/protobuf) v3

### Command

- **Compile for Rust**

  ```console
  make rust
  ```
  
- **Compile for Go**

  ```console
  make go
  ```
  
  The compilation will install `protoc`/ `protoc-gen-go` / `protoc-gen-go-grpc` locally.
  
- **Install protoc locally**

  ```console
  make install-protoc
  ```

  Then the `protoc` will install in `./bin/`.

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

### Go

Download the go module:

```console
go get github.com/GreptimeTeam/greptime-proto@main
```

Then use greptime-proto as the normal Go module:

```go
import (
        greptimev1 "github.com/GreptimeTeam/greptime-proto/go/greptime/v1"
)
...
```

## For SDK developers

GreptimeDB's gRPC service is built on top of [Arrow Flight RPC](https://arrow.apache.org/docs/format/Flight.html). You can find the Arrow's official implementation status of each programming language [here](https://arrow.apache.org/docs/status.html#flight-rpc).

> If you can't find the language you are using, you can always generate the Arrow Flight gRPC service from the raw protobuf definition [here](https://arrow.apache.org/docs/format/Flight.html#protocol-buffer-definitions). Or call into other language binding like C++. 

Once the Arrow Flight client is ready, you only need to care about the following 3 protobuf files to accomplish our SDK writing:

```console
.
├── greptime
│   └── v1
│       ├── column.proto        
│       ├── database.proto      
│       ├── ddl.proto            
```

> You can find all protobuf files in the directory "proto" under the project's root.

Right now the GreptimeDB only responds to Arrow Flight's `do_get` interface. All the reads and writes (that are from clients) are handled there. `do_get` needs a "ticket" as the input request, you need to serialize the `GreptimeRequest` message defined in "database.proto" for it.

There are 3 kinds of `GreptimeRequest`, which are:

- `InsertRequest`, carries the data to be ingested. It's a little verbose to assemble, especially the "column"s that define the schema and value of the input data. You can find the definition of "column" in "column.proto".
- `QueryRequest` has the SQL in it. Note that you can `INSERT INTO` GreptimeDB as well as `SELECT` it.
- `DdlRequest` defines the "Data Definition Language" request, like table creation or deletion. It's sometimes more representative than the normal SQL. `DdlRequest`s are defined in "ddl.proto".

There's also `RequestHeader` in the `GreptimeRequest` to specify the "catalog" and "schema" to be used in this request. If either one is not set (or left empty), GreptimeDB will use the default catalog "greptime" and default schema "public".

The response of `do_get` is handled in Arrow Flight's client. It's a stream of `FlightData`. You can find its definition in Arrow Flight's protobuf file. Special care must be taken when dealing with insertion, that GreptimeDB puts insertion result in `FlightData`'s metadata. Insertion results, either from `InsertRequest` or `INSERT INTO` SQL, are both have the same format, that "Affected Rows: x". "x" represents the rows that are successfully inserted. When dealing with this special `FlightData`, please ignore its `data_body` field, but directly strip the `app_metadata` field from it, and deserialize the bytes as `FlightMetadata` message. You will find the "affected rows" result in it.

We already have our SDK written in [Java](https://github.com/GreptimeTeam/greptimedb-client-java), [Rust]() and [Go](), feel free to take any of them as an example.
