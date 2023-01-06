package queue

import (
	"_models/ent"

	"go.uber.org/zap"
)

type QueueService struct {
	client *ent.Client
	logger *zap.Logger
}

func NewQueueService(client *ent.Client, logger *zap.Logger) *QueueService {
	return &QueueService{client: client, logger: logger}
}
