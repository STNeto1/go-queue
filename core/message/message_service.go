package message

import (
	"_core/queue"
	"_models/ent"
	"_models/ent/message"
	"context"
	"errors"

	"github.com/google/uuid"
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

func (m *MessageService) GetMessageFromUserQueue(ctx context.Context, tx *ent.Tx, user *ent.User, msgId uuid.UUID) (*ent.Message, error) {
	msg, err := tx.Message.
		Query().
		Where(message.IDEQ(msgId)).
		WithQueue(func(qq *ent.QueueQuery) {
			qq.WithUser()
		}).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	if msg.Edges.Queue.Edges.User.ID != user.ID {
		return nil, errors.New("user is not the owner of the queue")
	}

	return msg, nil
}
