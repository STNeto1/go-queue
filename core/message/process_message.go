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

type ProcessMessagePayload struct {
	MessageID uuid.UUID
	User      *ent.User
}

func (m *MessageService) ProcessMessage(ctx context.Context, payload *ProcessMessagePayload) (*ent.Message, error) {
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

	updatedMsg, err := tx.Message.
		UpdateOne(msg).
		SetStatus(schema.QueueMessageStatusProcessed).
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

	return updatedMsg, nil
}
