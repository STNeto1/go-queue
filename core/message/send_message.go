package message

import (
	"_core/queue"
	"_models/ent"
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type SendMessagePayload struct {
	QueueID       uuid.UUID
	User          *ent.User
	Body          string
	ContentType   string
	MaxRetries    uint
	SecondsOffset uint
}

func (m *MessageService) SendMessage(ctx context.Context, payload *SendMessagePayload) (*ent.Message, error) {
	q, err := m.qs.ShowQueue(ctx, queue.ShowQueuePayload{
		ID:   payload.QueueID,
		User: payload.User,
	})
	if err != nil {
		return nil, err
	}

	var retries uint = 5
	if payload.MaxRetries > 0 {
		retries = payload.MaxRetries
	}

	available := time.Now().Add(time.Duration(payload.SecondsOffset) * time.Second)

	msg, err := m.client.Message.Create().
		SetQueue(q).
		SetBody(payload.Body).
		SetContentType(payload.ContentType).
		SetMaxRetries(retries).
		SetAvailableFrom(available).
		Save(ctx)
	if err != nil {
		m.logger.Error("failed to create message", zap.Error(err))
		return nil, err
	}

	return msg, nil
}
