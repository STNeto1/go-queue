package message

import (
	"_models/ent"

	"go.uber.org/zap"
)

type MessageService struct {
	client *ent.Client
	logger *zap.Logger
}

func NewMessageService(client *ent.Client, logger *zap.Logger) *MessageService {
	return &MessageService{
		client: client,
		logger: logger,
	}
}
