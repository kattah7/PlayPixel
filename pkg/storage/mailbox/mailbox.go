package mailbox

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/PlayPixel/api/internal/db"
	"github.com/PlayPixel/api/internal/logger"
)

func AddMailbox(payload *MailboxAccount, ctx context.Context, pool db.Pool, log logger.Logger) error {
	if payload.SenderID == 0 {
		return fmt.Errorf("senderId is required")
	}

	if payload.SenderName == "" {
		return fmt.Errorf("senderName is required")
	}

	if payload.TargetID == 0 {
		return fmt.Errorf("targetId is required")
	}

	if payload.TargetName == "" {
		return fmt.Errorf("targetName is required")
	}

	if payload.SenderID == payload.TargetID {
		return fmt.Errorf("sender and target id cannot be the same")
	}

	var objOrArray struct{}
	if err := json.Unmarshal(payload.Item, &objOrArray); err != nil {
		return err
	}

	_, err := pool.Exec(ctx, `INSERT INTO mailbox (senderId, senderName, targetId, targetName, item) VALUES ($1, $2, $3, $4, $5)`, payload.SenderID, payload.SenderName, payload.TargetID, payload.TargetName, payload.Item)
	if err != nil {
		log.Errorw("failed to insert mailbox", "error", err)
		return err
	}

	return nil
}
