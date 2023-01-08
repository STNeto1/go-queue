package message_test

import (
	"_core/message"
	"_core/queue"
	"_models/ent"
	"_models/ent/enttest"
	"context"
	"testing"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func CreateMessageService(t *testing.T) (*message.MessageService, *ent.Client, *zap.Logger) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	logger := zap.NewNop()

	qs := queue.NewQueueService(client, logger)

	return message.NewMessageService(client, logger, qs), client, logger
}

func TestCreateMessageService(t *testing.T) {
	s, client, l := CreateMessageService(t)
	defer client.Close()
	defer l.Sync()

	assert.NotNil(t, s)
	assert.NotNil(t, client)
}

func TestErrorForNomessage(t *testing.T) {
	s, client, l := CreateMessageService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	q, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, q)
	assert.NoError(t, err)

	tx, err := client.Tx(context.Background())
	assert.NoError(t, err)

	foundMsg, err := s.GetMessageFromUserQueue(context.Background(), tx, usr, uuid.New())
	assert.Nil(t, foundMsg)
	assert.Error(t, err)

	err = tx.Commit()
	assert.NoError(t, err)
}

func TestErrorForUnauthorized(t *testing.T) {
	s, client, l := CreateMessageService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	usr2, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	q, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, q)
	assert.NoError(t, err)

	msg, err := createQueueMessage(client, usr, q, "some message")
	assert.NotNil(t, msg)
	assert.NoError(t, err)

	tx, err := client.Tx(context.Background())
	assert.NoError(t, err)

	foundMsg, err := s.GetMessageFromUserQueue(context.Background(), tx, usr2, msg.ID)
	assert.Nil(t, foundMsg)
	assert.Error(t, err)

	err = tx.Commit()
	assert.NoError(t, err)
}
