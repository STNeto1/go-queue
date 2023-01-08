package message

import (
	"_core/queue"
	"_models/ent"

	"go.uber.org/zap"
)

type MessageService struct {
	client *ent.Client
	logger *zap.Logger
	qs     *queue.QueueService
}

func NewMessageService(client *ent.Client, logger *zap.Logger, qs *queue.QueueService) *MessageService {
	return &MessageService{
		client: client,
		logger: logger,
		qs:     qs,
	}
}
