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

include!("../generated/greptime.v1.meta.rs");

mod mailbox;

use std::collections::HashMap;
use std::fmt::{Display, Formatter};
use std::hash::Hash;

pub const PROTOCOL_VERSION: u64 = 1;

#[derive(Default)]
pub struct PeerDict {
    peers: HashMap<Peer, usize>,
    index: usize,
}

impl PeerDict {
    pub fn get_or_insert(&mut self, peer: Peer) -> usize {
        let index = self.peers.entry(peer).or_insert_with(|| {
            let v = self.index;
            self.index += 1;
            v
        });

        *index
    }

    pub fn into_peers(self) -> Vec<Peer> {
        let mut array = vec![Peer::default(); self.index];
        for (p, i) in self.peers {
            array[i] = p;
        }
        array
    }
}

impl Display for Peer {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "peer-{}({})", self.id, self.addr)
    }
}

impl Peer {
    pub fn new(id: u64, addr: impl Into<String>) -> Self {
        Self {
            id,
            addr: addr.into(),
        }
    }
    pub fn empty(id: u64) -> Self {
        Self {
            id,
            addr: String::new(),
        }
    }
}

impl RequestHeader {
    #[inline]
    pub fn new(member_id: u64, role: Role, tracing_context: HashMap<String, String>) -> Self {
        Self {
            protocol_version: PROTOCOL_VERSION,
            member_id,
            role: role.into(),
            tracing_context,
        }
    }
}

impl ResponseHeader {
    #[inline]
    pub fn success() -> Self {
        Self {
            protocol_version: PROTOCOL_VERSION,
            ..Default::default()
        }
    }

    #[inline]
    pub fn failed(error: Error) -> Self {
        Self {
            protocol_version: PROTOCOL_VERSION,
            error: Some(error),
        }
    }

    #[inline]
    pub fn is_not_leader(&self) -> bool {
        if let Some(error) = &self.error {
            if error.code == ErrorCode::NotLeader as i32 {
                return true;
            }
        }
        false
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ErrorCode {
    NotEnoughAvailableDatanodes = 1,
    NotLeader = 2,
}

impl Error {
    #[inline]
    pub fn not_enough_available_datanodes(expected: usize, actual: usize) -> Self {
        Self {
            code: ErrorCode::NotEnoughAvailableDatanodes as i32,
            err_msg: format!(
                "There are not enough active datanodes, expected: {expected}, actual: {actual}."
            ),
        }
    }

    #[inline]
    pub fn is_not_leader() -> Self {
        Self {
            code: ErrorCode::NotLeader as i32,
            err_msg: "Current server is not leader".to_string(),
        }
    }
}

impl HeartbeatResponse {
    #[inline]
    pub fn is_not_leader(&self) -> bool {
        if let Some(header) = &self.header {
            return header.is_not_leader();
        }
        false
    }
}
macro_rules! gen_set_header {
    ($req: ty) => {
        impl $req {
            #[inline]
            pub fn set_header(
                &mut self,
                member_id: u64,
                role: Role,
                tracing_context: HashMap<String, String>,
            ) {
                match self.header.as_mut() {
                    Some(header) => {
                        header.member_id = member_id;
                        header.role = role.into();
                        header.tracing_context = tracing_context;
                    }
                    None => {
                        self.header = Some(RequestHeader::new(member_id, role, tracing_context));
                    }
                }
            }
        }
    };
}

gen_set_header!(HeartbeatRequest);
gen_set_header!(RangeRequest);
gen_set_header!(PutRequest);
gen_set_header!(BatchGetRequest);
gen_set_header!(BatchPutRequest);
gen_set_header!(BatchDeleteRequest);
gen_set_header!(CompareAndPutRequest);
gen_set_header!(DeleteRangeRequest);
gen_set_header!(LockRequest);
gen_set_header!(UnlockRequest);
gen_set_header!(DdlTaskRequest);
gen_set_header!(MigrateRegionRequest);
gen_set_header!(QueryProcedureRequest);
gen_set_header!(ProcedureDetailRequest);
gen_set_header!(ReconcileRequest);

#[cfg(test)]
mod tests {
    use std::vec;

    use prost::Message;
    use prost_types::field_descriptor_proto::{Label, Type};
    use prost_types::FileDescriptorSet;

    use super::*;

    #[derive(Clone, PartialEq, Message)]
    struct LegacyHeartbeatResponse {
        #[prost(message, optional, tag = "1")]
        header: Option<ResponseHeader>,
        #[prost(message, optional, tag = "2")]
        mailbox_message: Option<MailboxMessage>,
        #[prost(message, optional, tag = "3")]
        region_lease: Option<RegionLease>,
        #[prost(message, optional, tag = "4")]
        heartbeat_config: Option<HeartbeatConfig>,
    }

    #[test]
    fn test_heartbeat_response_extensions_round_trip() {
        let response = HeartbeatResponse {
            header: Some(ResponseHeader::success()),
            heartbeat_config: Some(HeartbeatConfig {
                heartbeat_interval_ms: 3_000,
                retry_interval_ms: 500,
                gc_enabled: true,
            }),
            extensions: HashMap::from([("test.extension".to_string(), vec![1, 2, 3])]),
            ..Default::default()
        };

        let decoded = HeartbeatResponse::decode(response.encode_to_vec().as_slice()).unwrap();
        assert_eq!(response, decoded);
    }

    #[test]
    fn test_heartbeat_response_extensions_wire_layout_is_stable() {
        let response = HeartbeatResponse {
            extensions: HashMap::from([("test.extension".to_string(), vec![1, 2, 3])]),
            ..Default::default()
        };
        let golden = [
            0x9a, 0x06, 0x15, 0x0a, 0x0e, b't', b'e', b's', b't', b'.', b'e', b'x', b't', b'e',
            b'n', b's', b'i', b'o', b'n', 0x12, 0x03, 0x01, 0x02, 0x03,
        ];

        assert_eq!(golden, response.encode_to_vec().as_slice());
        assert_eq!(
            response,
            HeartbeatResponse::decode(golden.as_slice()).unwrap()
        );
    }

    #[test]
    fn test_heartbeat_response_extensions_descriptor_is_stable() {
        let descriptor_set = FileDescriptorSet::decode(crate::v1::GREPTIME_GRPC_DESC).unwrap();
        let heartbeat_file = descriptor_set
            .file
            .iter()
            .find(|file| file.name.as_deref() == Some("greptime/v1/meta/heartbeat.proto"))
            .unwrap();
        let heartbeat_response = heartbeat_file
            .message_type
            .iter()
            .find(|message| message.name.as_deref() == Some("HeartbeatResponse"))
            .unwrap();
        let extensions = heartbeat_response
            .field
            .iter()
            .find(|field| field.name.as_deref() == Some("extensions"))
            .unwrap();

        assert_eq!(Some(99), extensions.number);
        assert_eq!(Some(Label::Repeated as i32), extensions.label);
        assert_eq!(Some(Type::Message as i32), extensions.r#type);
        assert_eq!(
            Some(".greptime.v1.meta.HeartbeatResponse.ExtensionsEntry"),
            extensions.type_name.as_deref()
        );

        let map_entry = heartbeat_response
            .nested_type
            .iter()
            .find(|message| message.name.as_deref() == Some("ExtensionsEntry"))
            .unwrap();
        assert_eq!(
            Some(true),
            map_entry
                .options
                .as_ref()
                .and_then(|options| options.map_entry)
        );

        let key = &map_entry.field[0];
        assert_eq!(
            (Some("key"), Some(1), Some(Type::String as i32)),
            (key.name.as_deref(), key.number, key.r#type)
        );
        let value = &map_entry.field[1];
        assert_eq!(
            (Some("value"), Some(2), Some(Type::Bytes as i32)),
            (value.name.as_deref(), value.number, value.r#type)
        );
    }

    #[test]
    fn test_heartbeat_response_extensions_wire_compatibility() {
        let new_response = HeartbeatResponse {
            header: Some(ResponseHeader::success()),
            heartbeat_config: Some(HeartbeatConfig {
                heartbeat_interval_ms: 3_000,
                retry_interval_ms: 500,
                gc_enabled: true,
            }),
            extensions: HashMap::from([("test.extension".to_string(), vec![1, 2, 3])]),
            ..Default::default()
        };
        let legacy_decoded =
            LegacyHeartbeatResponse::decode(new_response.encode_to_vec().as_slice()).unwrap();
        assert_eq!(new_response.header, legacy_decoded.header);
        assert_eq!(
            new_response.heartbeat_config,
            legacy_decoded.heartbeat_config
        );

        let legacy_response = LegacyHeartbeatResponse {
            header: Some(ResponseHeader::success()),
            mailbox_message: None,
            region_lease: None,
            heartbeat_config: Some(HeartbeatConfig {
                heartbeat_interval_ms: 4_000,
                retry_interval_ms: 1_000,
                gc_enabled: false,
            }),
        };
        let legacy_golden = [
            0x0a, 0x02, 0x08, 0x01, 0x22, 0x06, 0x08, 0xa0, 0x1f, 0x10, 0xe8, 0x07,
        ];
        assert_eq!(legacy_golden, legacy_response.encode_to_vec().as_slice());
        let new_decoded = HeartbeatResponse::decode(legacy_golden.as_slice()).unwrap();
        assert_eq!(legacy_response.header, new_decoded.header);
        assert_eq!(
            legacy_response.heartbeat_config,
            new_decoded.heartbeat_config
        );
        assert!(new_decoded.extensions.is_empty());
    }

    #[test]
    fn test_peer_dict() {
        let mut dict = PeerDict::default();

        dict.get_or_insert(Peer {
            id: 1,
            addr: "111".to_string(),
        });
        dict.get_or_insert(Peer {
            id: 2,
            addr: "222".to_string(),
        });
        dict.get_or_insert(Peer {
            id: 1,
            addr: "111".to_string(),
        });
        dict.get_or_insert(Peer {
            id: 1,
            addr: "111".to_string(),
        });
        dict.get_or_insert(Peer {
            id: 1,
            addr: "111".to_string(),
        });
        dict.get_or_insert(Peer {
            id: 1,
            addr: "111".to_string(),
        });
        dict.get_or_insert(Peer {
            id: 2,
            addr: "222".to_string(),
        });

        assert_eq!(2, dict.index);
        assert_eq!(
            vec![
                Peer {
                    id: 1,
                    addr: "111".to_string(),
                },
                Peer {
                    id: 2,
                    addr: "222".to_string(),
                }
            ],
            dict.into_peers()
        );
    }
}
