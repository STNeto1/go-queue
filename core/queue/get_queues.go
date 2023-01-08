package queue

import (
	"_models/ent"
	"_models/ent/queue"
	"_models/ent/user"
	"context"

	"go.uber.org/zap"
)

type GetQueuesPayload struct {
	Name string
	User *ent.User
}

func (q *QueueService) GetQueues(ctx context.Context, payload GetQueuesPayload) ([]*ent.Queue, error) {
	qb := q.client.Queue.Query().Where(queue.HasUserWith(user.ID(payload.User.ID)))

	if payload.Name != "" {
		qb = qb.Where(queue.NameContains(payload.Name))
	}

	queues, err := qb.All(ctx)
	if err != nil {
		q.logger.Error("failed to get queues", zap.Error(err))
		return nil, err
	}

	return queues, nil
}
