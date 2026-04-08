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

use std::env;
use std::error::Error;
use std::fs;
use std::path::PathBuf;

const LICENSE_HEADER: &str = r#"// Copyright 2023 Greptime Team
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

"#;

const PROTO_FILES: &[&str] = &[
    "proto/greptime/v1/database.proto",
    "proto/greptime/v1/health.proto",
    "proto/greptime/v1/wal.proto",
    "proto/greptime/v1/meta/common.proto",
    "proto/greptime/v1/meta/heartbeat.proto",
    "proto/greptime/v1/meta/config.proto",
    "proto/greptime/v1/meta/route.proto",
    "proto/greptime/v1/meta/ddl.proto",
    "proto/greptime/v1/meta/region.proto",
    "proto/greptime/v1/meta/store.proto",
    "proto/greptime/v1/meta/lock.proto",
    "proto/greptime/v1/meta/cluster.proto",
    "proto/greptime/v1/meta/procedure.proto",
    "proto/greptime/v1/region/server.proto",
    "proto/greptime/v1/flow/server.proto",
    "proto/greptime/v1/frontend/server.proto",
    "proto/greptime/v1/index/bloom_filter.proto",
    "proto/greptime/v1/index/inverted_index.proto",
    "proto/greptime/v1/index/vector_index.proto",
    "proto/prometheus/remote/remote.proto",
    "proto/substrait_extension/promql_plan.proto",
    "proto/substrait_extension/dist_plan.proto",
];

const GENERATED_FILES: &[&str] = &[
    "greptime.v1.rs",
    "greptime.v1.meta.rs",
    "greptime.v1.region.rs",
    "greptime.v1.flow.rs",
    "greptime.v1.frontend.rs",
    "greptime.v1.index.rs",
    "prometheus.rs",
    "substrait_extension.rs",
    "greptime_grpc_desc.bin",
];

fn main() -> Result<(), Box<dyn Error>> {
    run()
}

fn run() -> Result<(), Box<dyn Error>> {
    let mut args = env::args().skip(1);

    match args.next().as_deref() {
        Some("generate-rust") if args.next().is_none() => generate_rust(),
        _ => Err("usage: cargo run --manifest-path xtask/Cargo.toml -- generate-rust".into()),
    }
}

fn generate_rust() -> Result<(), Box<dyn Error>> {
    let repo_root = PathBuf::from(env!("CARGO_MANIFEST_DIR"))
        .parent()
        .ok_or("xtask manifest must live under the repository root")?
        .to_path_buf();
    let generated_dir = repo_root.join("src/generated");

    fs::create_dir_all(&generated_dir)?;
    for generated_file in GENERATED_FILES {
        let path = generated_dir.join(generated_file);
        if path.exists() {
            fs::remove_file(path)?;
        }
    }

    let proto_files: Vec<PathBuf> = PROTO_FILES
        .iter()
        .map(|path| repo_root.join(path))
        .collect();
    let include_dirs = [repo_root.join("proto")];

    tonic_prost_build::configure()
        .out_dir(&generated_dir)
        .file_descriptor_set_path(generated_dir.join("greptime_grpc_desc.bin"))
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
        .bytes(".greptime.v1.ArrowIpc")
        .compile_protos(&proto_files, &include_dirs)?;

    for generated_file in GENERATED_FILES {
        if generated_file.ends_with(".rs") {
            prepend_license_header(&generated_dir.join(generated_file))?;
        }
    }

    Ok(())
}

fn prepend_license_header(path: &PathBuf) -> Result<(), Box<dyn Error>> {
    let contents = fs::read_to_string(path)?;
    if contents.starts_with(LICENSE_HEADER) {
        return Ok(());
    }

    fs::write(path, format!("{LICENSE_HEADER}{contents}"))?;
    Ok(())
}
