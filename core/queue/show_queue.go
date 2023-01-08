package queue

import (
	"_models/ent"
	"_models/ent/queue"
	"_models/ent/user"
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ShowQueuePayload struct {
	ID   uuid.UUID
	User *ent.User
}

func (q *QueueService) ShowQueue(ctx context.Context, payload ShowQueuePayload) (*ent.Queue, error) {
	queue, err := q.client.Queue.Query().Where(queue.HasUserWith(user.ID(payload.User.ID))).Where(queue.ID(payload.ID)).First(ctx)
	if err != nil {
		q.logger.Error("failed to show queue", zap.Error(err))
		return nil, err
	}

	return queue, nil
}
