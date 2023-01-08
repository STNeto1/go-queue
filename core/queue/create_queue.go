package queue

import (
	"_models/ent"
	"context"
	"errors"

	"go.uber.org/zap"
)

type CreateQueuePayload struct {
	Name string
	User *ent.User
}

func (q *QueueService) CreateQueue(ctx context.Context, payload CreateQueuePayload) (*ent.Queue, error) {
	queue, err := q.client.Queue.Create().SetName(payload.Name).AddUser(payload.User).Save(ctx)
	if err != nil {
		q.logger.Error("error while creating queue", zap.Error(err))
		return nil, errors.New("error while creating queue")
	}

	return queue, nil
}
