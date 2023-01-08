package message_test

import (
	"_core/message"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteMessageFromQueue(t *testing.T) {
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

	err = s.DeleteMessage(context.Background(), &message.DeleteMessagePayload{
		MessageID: msg.ID,
		User:      usr,
	})
	assert.NoError(t, err)
}
