package mailbox

import (
	"encoding/json"
	"time"
)

type MailboxAccount struct {
	Payload    string          `json:"payload"`
	SenderID   int64           `json:"senderId"`
	SenderName string          `json:"senderName"`
	TargetID   int64           `json:"targetId"`
	TargetName string          `json:"targetName"`
	Item       json.RawMessage `json:"item"`
	Created    time.Time       `json:"created"`
	Message    string          `json:"message"`
}
