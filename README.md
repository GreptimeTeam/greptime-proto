# greptime-proto

GreptimeDB protobuf definitions and pre-generated Rust bindings.

## Build

### Requirement

- Rust consumers do not need `protoc`.
- Maintainers need [google/protobuf][protobuf] v3 to regenerate Rust bindings after `.proto` changes.

### Command

- **Generate Rust bindings**

  ```console
  make rust
  ```

- **Compile for Go**

  ```console
  make go
  ```

- **Compile for Java**

  ```console
  make java
  ```

  The compilation for Go and Java will use builder container `namely/protoc-all`.

## Usage

### Rust

```toml
greptime-proto = "0.1"
```

```rust
// To use the GreptimeDB's gRPC service:
use greptime_proto::v1::*;

// To talk to GreptimeDB Meta service:
use greptime_proto::v1::meta::*;

// To request GreptimeDB as Prometheus remote read/write:
use greptime_proto::prometheus::remote::*;
```

When working in this repository, regenerate the checked-in Rust bindings after `.proto` changes with:

```console
cargo run --manifest-path xtask/Cargo.toml -- generate-rust
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

GreptimeDB's gRPC service is built on top of [Arrow Flight RPC][flight].  You can find the Arrow's
official implementation status of each programming language [flight rpc][flight-rpc].

> If you can't find the language you are using, you can always generate the Arrow Flight gRPC
> service from the raw protobuf definition [flight protobuf definitions][flight-protobuf]. Or call
> into other language binding like C++.

Once the Arrow Flight client is ready, you only need to care about the following protobuf files to
accomplish our SDK writing:

```console
.
├── greptime
│   └── v1
│       ├── column.proto
│       ├── common.proto
│       ├── database.proto
│       ├── ddl.proto
│       ├── health.proto
│       └── prom.proto
```

> You can find all protobuf files in the directory "proto" under the project's root.

Right now the GreptimeDB only responds to Arrow Flight's `do_get` interface. All the reads and
writes (that are from clients) are handled there. `do_get` needs a "ticket" as the input request,
you need to serialize the `GreptimeRequest` message defined in "database.proto" for it.

There are 3 kinds of `GreptimeRequest`, which are:

- `InsertRequest`, carries the data to be ingested. It's a little verbose to assemble, especially
  the "column"s that define the schema and value of the input data. You can find the definition of
  "column" in "column.proto".
- `QueryRequest` has the SQL in it. Note that you can `INSERT INTO` GreptimeDB as well as `SELECT`
  it.
- `DdlRequest` defines the "Data Definition Language" request, like table creation or deletion. It's
  sometimes more representative than the normal SQL. `DdlRequest`s are defined in "ddl.proto".

There's also `RequestHeader` in the `GreptimeRequest` to specify the "catalog" and "schema" to be
used in this request. If either one is not set (or left empty), GreptimeDB will use the default
catalog "greptime" and default schema "public".

The response of `do_get` is handled in Arrow Flight's client. It's a stream of `FlightData`. You can
find its definition in Arrow Flight's protobuf file. Special care must be taken when dealing with
insertion, that GreptimeDB puts insertion result in `FlightData`'s metadata. Insertion results,
either from `InsertRequest` or `INSERT INTO` SQL, are both have the same format, that "Affected
Rows: x". "x" represents the rows that are successfully inserted. When dealing with this special
`FlightData`, please ignore its `data_body` field, but directly strip the `app_metadata` field from
it, and deserialize the bytes as `FlightMetadata` message. You will find the "affected rows" result
in it.

We already have our SDK written in [Java][java-sdk], [Rust][rust-sdk] and [Go][go-sdk], feel free to take
any of them as an example.

<!-- links -->
[protobuf]: https://github.com/protocolbuffers/protobuf
[flight]: https://arrow.apache.org/docs/format/Flight.html
[flight-rpc]: https://arrow.apache.org/docs/status.html#flight-rpc
[flight-protobuf]: https://arrow.apache.org/docs/format/Flight.html#protocol-buffer-definitions
[java-sdk]: https://github.com/GreptimeTeam/greptimedb-ingester-java
[rust-sdk]: https://github.com/GreptimeTeam/greptimedb-ingester-rust
[go-sdk]: https://github.com/GreptimeTeam/greptimedb-ingester-go
