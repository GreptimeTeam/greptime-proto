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
        .compile(
            &[
                "proto/greptime/v1/database.proto",
                "proto/greptime/v1/health.proto",
                "proto/greptime/v1/meta/common.proto",
                "proto/greptime/v1/meta/heartbeat.proto",
                "proto/greptime/v1/meta/route.proto",
                "proto/greptime/v1/meta/store.proto",
                "proto/greptime/v1/meta/lock.proto",
                "proto/greptime/v1/meta/cluster.proto",
                "proto/prometheus/remote/remote.proto",
            ],
            &["proto"],
        )
        .expect("compile proto");
}
