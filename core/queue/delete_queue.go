package queue

import (
	"_models/ent"
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type DeleteQueuePayload struct {
	ID   uuid.UUID
	User *ent.User
}

func (q *QueueService) DeleteQueue(ctx context.Context, payload DeleteQueuePayload) error {
	queue, err := q.ShowQueue(ctx, ShowQueuePayload{
		ID:   payload.ID,
		User: payload.User,
	})
	if err != nil {
		q.logger.Error("queue not found", zap.Error(err))
		return err
	}

	err = q.client.Queue.DeleteOne(queue).Exec(ctx)
	if err != nil {
		q.logger.Error("failed to delete queue", zap.Error(err))
		return err
	}

	return nil
}
