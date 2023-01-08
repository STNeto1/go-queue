package message_test

import (
	"_core/message"
	"_models/ent/schema"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessMessageFromQueue(t *testing.T) {
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

	updated, err := s.ProcessMessage(context.Background(), &message.ProcessMessagePayload{
		MessageID: msg.ID,
		User:      usr,
	})
	assert.NoError(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, schema.QueueMessageStatusProcessed, updated.Status)
}
