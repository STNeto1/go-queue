package message

import (
	lib "_lib"
	"_models/ent"
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type DeleteMessagePayload struct {
	MessageID uuid.UUID
	User      *ent.User
}

func (m *MessageService) DeleteMessage(ctx context.Context, payload *DeleteMessagePayload) error {
	tx, err := m.client.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		m.logger.Error("failed to create transaction", zap.Error(err))
		return err
	}

	msg, err := m.GetMessageFromUserQueue(ctx, tx, payload.User, payload.MessageID)
	if err != nil {
		m.logger.Error("failed to get message", zap.Error(err))
		return lib.Rollback(tx, errors.New("failed to get message"))
	}

	err = tx.Message.
		DeleteOne(msg).
		Exec(ctx)
	if err != nil {
		m.logger.Error("failed to delete message", zap.Error(err))
		return lib.Rollback(tx, errors.New("failed to delete message"))
	}

	err = tx.Commit()
	if err != nil {
		m.logger.Error("failed to commit transaction", zap.Error(err))
		return errors.New("failed to commit transaction")
	}

	return nil
}
