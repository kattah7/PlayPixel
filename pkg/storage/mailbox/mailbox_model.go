package mailbox

import (
	"encoding/json"
	"time"
)

type Mailbox struct {
	Payload string `json:"payload"`

	// read
	RobloxID int64 `json:"robloxId"`

	// insert
	SenderID   int64           `json:"senderId"`
	SenderName string          `json:"senderName"`
	TargetID   int64           `json:"targetId"`
	TargetName string          `json:"targetName"`
	Item       json.RawMessage `json:"item"`
	Created    time.Time       `json:"created"`
	Message    string          `json:"message"`

	// delete
	UID              int64 `json:"uid"`
	UIDMatchRobloxId int64 `json:"uidMatchRobloxId"`
}

type ReadMailboxResponse struct {
	UID        int64           `json:"uid"`
	SenderID   int64           `json:"senderId"`
	SenderName string          `json:"senderName"`
	Item       json.RawMessage `json:"item"`
	Created    time.Time       `json:"mail_created"`
	Message    string          `json:"message"`
}
