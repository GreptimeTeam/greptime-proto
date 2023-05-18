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

use crate::v1::meta::mailbox_message::Payload;
use crate::v1::meta::MailboxMessage;
use serde::Serialize;
use serde_json::Result;
use std::fmt::{Display, Formatter};

impl MailboxMessage {
    pub fn json_message<T>(
        subject: &str,
        from: &str,
        to: &str,
        timestamp_millis: i64,
        payload: &T,
    ) -> Result<MailboxMessage>
    where
        T: ?Sized + Serialize + Display,
    {
        let payload = serde_json::to_string(payload)?;
        Ok(MailboxMessage {
            id: 0, // "id" will be set by the mailbox.
            subject: subject.to_string(),
            from: from.to_string(),
            to: to.to_string(),
            timestamp_millis,
            payload: Some(Payload::Json(payload)),
        })
    }
}

impl Display for MailboxMessage {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "MailboxMessage(id={}, subject={}, from={}, to={}, timestamp_millis={}, payload={:?})",
            self.id, self.subject, self.from, self.to, self.timestamp_millis, self.payload
        )
    }
}
