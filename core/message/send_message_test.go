package message_test

import (
	"_core/message"
	"_models/ent"
	"_models/ent/schema"
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func createUser(client *ent.Client) (*ent.User, error) {
	return client.User.
		Create().
		SetName("John Doe").
		SetEmail(fmt.Sprintf("some-mail-%d@mail.com", rand.Int())).
		SetPassword("some-password").
		Save(context.Background())
}

func createQueue(client *ent.Client, usr *ent.User, name string) (*ent.Queue, error) {
	return client.Queue.
		Create().
		SetName(name).
		SetUser(usr).
		Save(context.Background())
}

func TestSendMessageToQueue(t *testing.T) {
	s, client, l := CreateMessageService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	q, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, q)
	assert.NoError(t, err)

	msg, err := s.SendMessage(context.Background(), &message.SendMessagePayload{
		QueueRef:      q.Ref,
		User:          usr,
		Body:          "some-body",
		ContentType:   "some-content-type",
		MaxRetries:    5,
		SecondsOffset: 0,
	})
	assert.NotNil(t, msg)
	assert.NoError(t, err)
}

func TestSendMessageToQueueDefaultValues(t *testing.T) {
	s, client, l := CreateMessageService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	q, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, q)
	assert.NoError(t, err)

	msg, err := s.SendMessage(context.Background(), &message.SendMessagePayload{
		QueueRef:      q.Ref,
		User:          usr,
		Body:          "some-body",
		ContentType:   "some-content-type",
		MaxRetries:    0,
		SecondsOffset: 0,
	})
	assert.NotNil(t, msg)
	assert.NoError(t, err)
	assert.Equal(t, uint(5), msg.MaxRetries)
	assert.Equal(t, schema.QueueMessageStatusPending, msg.Status)
}

func TestSendMessageToQueueWithOffset(t *testing.T) {
	s, client, l := CreateMessageService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	q, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, q)
	assert.NoError(t, err)

	msg, err := s.SendMessage(context.Background(), &message.SendMessagePayload{
		QueueRef:      q.Ref,
		User:          usr,
		Body:          "some-body",
		ContentType:   "some-content-type",
		MaxRetries:    0,
		SecondsOffset: 10,
	})
	assert.NotNil(t, msg)
	assert.NoError(t, err)
	assert.Greater(t, msg.AvailableFrom.Unix(), time.Now().Unix())
}
