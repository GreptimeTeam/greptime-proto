// Copyright 2023 Greptime Team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

use std::path::PathBuf;

fn main() {
    let out_dir = PathBuf::from(
        std::env::var("OUT_DIR")
            .expect("cargo built-in env value 'OUT_DIR' must be set during compilation"),
    );

    tonic_build::configure()
        .file_descriptor_set_path(out_dir.join("greptime_grpc_desc.bin"))
        .type_attribute(
            ".greptime.v1.SemanticType",
            "#[derive(::serde::Serialize, ::serde::Deserialize)]",
        )
        .type_attribute(
            ".greptime.v1.meta.Peer",
            "#[derive(::serde::Serialize, ::serde::Deserialize)]",
        )
        .type_attribute(
            ".greptime.v1.meta.HeartbeatRequest.node_workloads",
            "#[derive(::serde::Serialize, ::serde::Deserialize)]",
        )
        .type_attribute(
            ".greptime.v1.meta.DatanodeWorkloads",
            "#[derive(::serde::Serialize, ::serde::Deserialize)]",
        )
        .type_attribute(
            ".greptime.v1.meta.FrontendWorkloads",
            "#[derive(::serde::Serialize, ::serde::Deserialize)]",
        )
        .type_attribute(
            ".greptime.v1.meta.FlownodeWorkloads",
            "#[derive(::serde::Serialize, ::serde::Deserialize)]",
        )
        .enum_attribute(
            "region.RegionRequest.body",
            "#[derive(strum_macros::AsRefStr)]",
        )
        .bytes([".greptime.v1.ArrowIpc"])
        .compile_protos(
            &[
                "proto/greptime/v1/database.proto",
                "proto/greptime/v1/health.proto",
                "proto/greptime/v1/wal.proto",
                "proto/greptime/v1/meta/common.proto",
                "proto/greptime/v1/meta/heartbeat.proto",
                "proto/greptime/v1/meta/route.proto",
                "proto/greptime/v1/meta/ddl.proto",
                "proto/greptime/v1/meta/region.proto",
                "proto/greptime/v1/meta/store.proto",
                "proto/greptime/v1/meta/lock.proto",
                "proto/greptime/v1/meta/cluster.proto",
                "proto/greptime/v1/meta/procedure.proto",
                "proto/greptime/v1/region/server.proto",
                "proto/greptime/v1/flow/server.proto",
                "proto/greptime/v1/index/bloom_filter.proto",
                "proto/greptime/v1/index/inverted_index.proto",
                "proto/prometheus/remote/remote.proto",
                "proto/substrait_extension/promql_plan.proto",
                "proto/substrait_extension/dist_plan.proto",
            ],
            &["proto"],
        )
        .expect("compile proto");
}
