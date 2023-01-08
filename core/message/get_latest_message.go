package message

import (
	cq "_core/queue"
	lib "_lib"
	"_models/ent"
	"_models/ent/message"
	"_models/ent/queue"
	"_models/ent/schema"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type GetLatestMessagePayload struct {
	QueueRef uuid.UUID
	User     *ent.User
}

func (m *MessageService) GetLatestMessage(ctx context.Context, payload *GetLatestMessagePayload) (*ent.Message, error) {
	q, err := m.qs.ShowQueueFromRef(ctx, cq.ShowQueuePayload{
		ID:   payload.QueueRef,
		User: payload.User,
	})
	if err != nil {
		return nil, err
	}

	tx, err := m.client.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		m.logger.Error("failed to create transaction", zap.Error(err))
		return nil, err
	}

	msg, err := tx.Message.
		Query().
		Where(message.HasQueueWith(queue.IDEQ(q.ID))).
		Where(message.Status(schema.QueueMessageStatusPending)).
		Where(message.AvailableFromLTE(time.Now())).
		Order(ent.Desc(message.FieldAvailableFrom)).
		First(ctx)
	if err != nil {
		m.logger.Error("failed to get latest message", zap.Error(err))
		return nil, lib.Rollback(tx, errors.New("failed to get latest message"))
	}

	updatedMsg, err := tx.Message.
		UpdateOne(msg).
		SetStatus(schema.QueueMessageStatusProcessing).
		Save(ctx)
	if err != nil {
		m.logger.Error("failed to update message status", zap.Error(err))
		return nil, lib.Rollback(tx, errors.New("failed to update message status"))
	}

	err = tx.Commit()
	if err != nil {
		m.logger.Error("failed to commit transaction", zap.Error(err))
		return nil, err
	}

	return updatedMsg, nil
}
