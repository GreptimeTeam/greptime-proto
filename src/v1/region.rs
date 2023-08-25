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

tonic::include_proto!("greptime.v1.region");

#[cfg(test)]
mod test {
    use crate::v1::region::region_request::Body as RegionRequest;
    use crate::v1::region::InsertRequests;

    #[test]
    fn test_region_request_name() {
        let request = RegionRequest::Inserts(InsertRequests { requests: vec![] });
        assert_eq!("Inserts", request.as_ref());
    }
}
