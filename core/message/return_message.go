package message

import (
	lib "_lib"
	"_models/ent"
	"_models/ent/schema"
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ReturnMessagePayload struct {
	MessageID uuid.UUID
	User      *ent.User
}

func (m *MessageService) ReturnMessage(ctx context.Context, payload *ReturnMessagePayload) (*ent.Message, error) {
	tx, err := m.client.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		m.logger.Error("failed to create transaction", zap.Error(err))
		return nil, err
	}

	msg, err := m.GetMessageFromUserQueue(ctx, tx, payload.User, payload.MessageID)
	if err != nil {
		m.logger.Error("failed to get message", zap.Error(err))
		return nil, lib.Rollback(tx, errors.New("failed to get message"))
	}

	status := schema.QueueMessageStatusProcessing
	if msg.MaxRetries == msg.Retries+1 {
		status = schema.QueueMessageStatusFailed
	}

	updated, err := tx.Message.
		UpdateOne(msg).
		SetStatus(status).
		SetRetries(msg.Retries + 1).
		Save(ctx)
	if err != nil {
		m.logger.Error("failed to update message status", zap.Error(err))
		return nil, lib.Rollback(tx, errors.New("failed to update message status"))
	}

	err = tx.Commit()
	if err != nil {
		m.logger.Error("failed to commit transaction", zap.Error(err))
		return nil, errors.New("failed to commit transaction")
	}

	return updated, nil
}
