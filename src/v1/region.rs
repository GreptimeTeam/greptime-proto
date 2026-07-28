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

include!("../generated/greptime.v1.region.rs");

#[cfg(test)]
mod test {
    use prost::Message;

    use crate::v1::region::region_request::Body as RegionRequest;
    use crate::v1::region::{CloseRequest, CompactRequest, CompactionTimeRange, InsertRequests};
    use crate::v1::TimeUnit;

    #[derive(Clone, PartialEq, Message)]
    struct LegacyCompactRequest {
        #[prost(uint64, tag = "1")]
        region_id: u64,
        #[prost(uint32, tag = "4")]
        parallelism: u32,
    }

    #[test]
    fn test_region_request_name() {
        let request = RegionRequest::Inserts(InsertRequests { requests: vec![] });
        assert_eq!("Inserts", request.as_ref());
    }

    #[test]
    fn test_compact_request_time_range_round_trip() {
        let request = CompactRequest {
            region_id: 42,
            time_range: Some(CompactionTimeRange {
                start: 1_767_225_600_000,
                end: 1_769_904_000_000,
                time_unit: TimeUnit::Millisecond as i32,
            }),
            ..Default::default()
        };

        let decoded = CompactRequest::decode(request.encode_to_vec().as_slice()).unwrap();
        assert_eq!(request, decoded);
    }

    #[test]
    fn test_compact_request_time_range_wire_compatibility() {
        let new_request = CompactRequest {
            region_id: 42,
            parallelism: 4,
            time_range: Some(CompactionTimeRange {
                start: 1_767_225_600,
                end: 1_769_904_000,
                time_unit: TimeUnit::Second as i32,
            }),
            ..Default::default()
        };
        let legacy_decoded =
            LegacyCompactRequest::decode(new_request.encode_to_vec().as_slice()).unwrap();
        assert_eq!(42, legacy_decoded.region_id);
        assert_eq!(4, legacy_decoded.parallelism);

        let legacy_request = LegacyCompactRequest {
            region_id: 43,
            parallelism: 2,
        };
        let new_decoded =
            CompactRequest::decode(legacy_request.encode_to_vec().as_slice()).unwrap();
        assert_eq!(43, new_decoded.region_id);
        assert_eq!(2, new_decoded.parallelism);
        assert!(new_decoded.time_range.is_none());
    }

    #[test]
    fn test_close_request_flush_on_close() {
        let request = CloseRequest {
            region_id: 42,
            flush_on_close: true,
        };

        assert!(request.flush_on_close);
    }
}
