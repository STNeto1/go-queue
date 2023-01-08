package message

import (
	"_core/queue"
	lib "_lib"
	"_models/ent"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type BulkMessageElement struct {
	Body          string
	ContentType   string
	MaxRetries    uint
	SecondsOffset uint
}

type SendBulkMessagesPayload struct {
	QueueID  uuid.UUID
	User     *ent.User
	Messages []BulkMessageElement
}

func (m *MessageService) SendBulkMessages(ctx context.Context, payload *SendBulkMessagesPayload) ([]*ent.Message, error) {
	q, err := m.qs.ShowQueue(ctx, queue.ShowQueuePayload{
		ID:   payload.QueueID,
		User: payload.User,
	})
	if err != nil {
		return nil, err
	}

	tx, err := m.client.Tx(ctx)
	if err != nil {
		m.logger.Error("failed to create transaction", zap.Error(err))
		return nil, err
	}

	var messages []*ent.Message
	for _, msg := range payload.Messages {
		var retries uint = 5
		if msg.MaxRetries > 0 {
			retries = msg.MaxRetries
		}

		available := time.Now().Add(time.Duration(msg.SecondsOffset) * time.Second)

		newMsg, err := tx.Message.
			Create().
			SetQueue(q).
			SetBody(msg.Body).
			SetContentType(msg.ContentType).
			SetMaxRetries(retries).
			SetAvailableFrom(available).
			Save(ctx)
		if err != nil {
			m.logger.Error("failed to create message", zap.Error(err))
			return nil, lib.Rollback(tx, errors.New("failed to create message"))
		}

		messages = append(messages, newMsg)
	}

	if err := tx.Commit(); err != nil {
		m.logger.Error("failed to commit transaction", zap.Error(err))
		return nil, err
	}

	return messages, nil
}
