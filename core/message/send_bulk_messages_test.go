package message_test

import (
	"_core/message"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendBulkMessages(t *testing.T) {
	s, client, l := CreateMessageService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	q, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, q)
	assert.NoError(t, err)

	msgs, err := s.SendBulkMessages(context.Background(), &message.SendBulkMessagesPayload{
		QueueID: q.ID,
		User:    usr,
		Messages: []message.BulkMessageElement{
			{
				Body:          "some-body",
				ContentType:   "some-content-type",
				MaxRetries:    5,
				SecondsOffset: 0,
			},
		}})
	assert.NotNil(t, msgs)
	assert.NoError(t, err)
	assert.Len(t, msgs, 1)
}
