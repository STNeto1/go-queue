package message_test

import (
	"_core/message"
	"_models/ent/schema"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReturnErrorWhenUnauthorizedQueueAccess(t *testing.T) {

}

func TestReturnMessageToQueue(t *testing.T) {
	s, client, l := CreateMessageService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	q, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, q)
	assert.NoError(t, err)

	msg, err := createQueueMessage(client, usr, q, "some message")
	assert.NotNil(t, msg)
	assert.NoError(t, err)

	updatedMsg, err := s.ReturnMessage(context.Background(), &message.ReturnMessagePayload{
		MessageID: msg.ID,
		User:      usr,
	})
	assert.NoError(t, err)
	assert.NotNil(t, updatedMsg)
	assert.Equal(t, schema.QueueMessageStatusProcessing, updatedMsg.Status)
	assert.Equal(t, uint(1), updatedMsg.Retries)
}

func TestReturnMessageToMaxRetries(t *testing.T) {
	s, client, l := CreateMessageService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	q, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, q)
	assert.NoError(t, err)

	msg, err := client.Message.Create().
		SetQueue(q).
		SetBody("some body").
		SetContentType("any content type").
		SetMaxRetries(1).
		SetAvailableFrom(time.Now()).
		Save(context.Background())
	assert.NotNil(t, msg)
	assert.NoError(t, err)

	updatedMsg, err := s.ReturnMessage(context.Background(), &message.ReturnMessagePayload{
		MessageID: msg.ID,
		User:      usr,
	})
	assert.NoError(t, err)
	assert.NotNil(t, updatedMsg)
	assert.Equal(t, schema.QueueMessageStatusFailed, updatedMsg.Status)
	assert.Equal(t, uint(1), updatedMsg.Retries)
}
