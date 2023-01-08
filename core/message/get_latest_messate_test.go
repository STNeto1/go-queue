package message_test

import (
	"_core/message"
	"_models/ent"
	"_models/ent/schema"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func createQueueMessage(client *ent.Client, usr *ent.User, q *ent.Queue, body string) (*ent.Message, error) {
	return client.Message.Create().
		SetQueue(q).
		SetBody(body).
		SetContentType("any content type").
		SetMaxRetries(5).
		SetAvailableFrom(time.Now()).
		Save(context.Background())
}

func TestErrorWhenReadingQueueWithNoMessages(t *testing.T) {
	s, client, l := CreateMessageService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	q, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, q)
	assert.NoError(t, err)

	msg, err := s.GetLatestMessage(context.Background(), &message.GetLatestMessagePayload{
		QueueID: q.ID,
		User:    usr,
	})
	assert.Nil(t, msg)
	assert.Error(t, err)
}

func TestGetLatestMessage(t *testing.T) {
	s, client, l := CreateMessageService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	q, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, q)
	assert.NoError(t, err)

	msg1, err := createQueueMessage(client, usr, q, "some message")
	assert.NotNil(t, msg1)
	assert.NoError(t, err)

	msg2, err := s.GetLatestMessage(context.Background(), &message.GetLatestMessagePayload{
		QueueID: q.ID,
		User:    usr,
	})
	assert.NotNil(t, msg2)
	assert.NoError(t, err)
	assert.Equal(t, msg1.ID, msg2.ID)
	assert.Equal(t, schema.QueueMessageStatusProcessing, msg2.Status)

}
