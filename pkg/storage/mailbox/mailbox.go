package mailbox

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/PlayPixel/api/internal/db"
	"github.com/PlayPixel/api/internal/logger"
)

const LIMIT = 100

func AddMailbox(payload *Mailbox, ctx context.Context, pool db.Pool, log logger.Logger) error {
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

	var count int
	err := pool.QueryRow(ctx, `SELECT COUNT(*) FROM mailbox WHERE targetId = $1`, payload.TargetID).Scan(&count)
	if err != nil {
		log.Errorw("failed to count mailbox", "error", err)
		return err
	}

	if count >= LIMIT {
		return fmt.Errorf("FULL")
	}

	if payload.Message == "" {
		payload.Message = "You have a new item in your mailbox!"
	}

	_, err2 := pool.Exec(ctx,
		`INSERT INTO mailbox (senderId, senderName, targetId, targetName, item, message)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		payload.SenderID, payload.SenderName, payload.TargetID, payload.TargetName, payload.Item, payload.Message,
	)

	if err2 != nil {
		log.Errorw("failed to insert mailbox", "error", err2)
		return err2
	}

	return nil
}

func ReadMailbox(payload *Mailbox, ctx context.Context, pool db.Pool, log logger.Logger) ([]ReadMailboxResponse, error) {
	if payload.RobloxID == 0 {
		return nil, fmt.Errorf("robloxId is required")
	}

	query := `SELECT id, senderId, senderName, item, created, message FROM mailbox WHERE targetId = $1`
	rows, err := pool.Query(ctx, query, payload.RobloxID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mailboxes []ReadMailboxResponse
	for rows.Next() {
		var UniqueID int64
		var SenderID int64
		var SenderName string
		var Item json.RawMessage
		var Created time.Time
		var Message string

		if err := rows.Scan(&UniqueID, &SenderID, &SenderName, &Item, &Created, &Message); err != nil {
			log.Errorw("failed to scan mailbox", "error", err)
			return nil, err
		}

		// Append each mailbox to the slice
		mailboxes = append(mailboxes, ReadMailboxResponse{
			UID:        UniqueID,
			SenderID:   SenderID,
			SenderName: SenderName,
			Item:       Item,
			Created:    Created,
			Message:    Message,
		})
	}

	if mailboxes == nil {
		return nil, fmt.Errorf("no mail found for robloxId %d", payload.RobloxID)
	}

	return mailboxes, nil
}

func DeleteMailbox(payload *Mailbox, ctx context.Context, pool db.Pool, log logger.Logger) error {
	if payload.UID == 0 {
		return fmt.Errorf("uid is required")
	}

	if payload.UIDMatchRobloxId == 0 {
		return fmt.Errorf("uidMatchRobloxId is required")
	}

	result, err := pool.Exec(ctx, `DELETE FROM mailbox WHERE id = $1 AND targetId = $2`, payload.UID, payload.UIDMatchRobloxId)
	if err != nil {
		log.Errorw("failed to delete mailbox", "error", err)
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were affected by the delete operation")
	}

	return nil
}
