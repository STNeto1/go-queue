package queue_test

import (
	qm "_core/queue"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestShowQueue(t *testing.T) {
	s, client, l := CreateQueueService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	queue, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, queue)
	assert.NoError(t, err)

	q, err := s.ShowQueue(context.Background(), qm.ShowQueuePayload{
		ID:   queue.ID,
		User: usr,
	})

	assert.NotNil(t, q)
	assert.NoError(t, err)
}

func TestShowQueueWithInvalidID(t *testing.T) {
	s, client, l := CreateQueueService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	queue, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, queue)
	assert.NoError(t, err)

	q, err := s.ShowQueue(context.Background(), qm.ShowQueuePayload{
		ID:   uuid.New(),
		User: usr,
	})

	assert.Nil(t, q)
	assert.Error(t, err)
}

func TestShowQueueWithInvalidUser(t *testing.T) {
	s, client, l := CreateQueueService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	usr2, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	queue, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, queue)
	assert.NoError(t, err)

	q, err := s.ShowQueue(context.Background(), qm.ShowQueuePayload{
		ID:   queue.ID,
		User: usr2,
	})

	assert.Nil(t, q)
	assert.Error(t, err)
}

func TestShowQueueFromRef(t *testing.T) {
	s, client, l := CreateQueueService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	queue, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, queue)
	assert.NoError(t, err)

	q, err := s.ShowQueueFromRef(context.Background(), qm.ShowQueuePayload{
		ID:   queue.Ref,
		User: usr,
	})

	assert.NotNil(t, q)
	assert.NoError(t, err)
}

func TestShowQueueWithInvalidRef(t *testing.T) {
	s, client, l := CreateQueueService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	queue, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, queue)
	assert.NoError(t, err)

	q, err := s.ShowQueueFromRef(context.Background(), qm.ShowQueuePayload{
		ID:   uuid.New(),
		User: usr,
	})

	assert.Nil(t, q)
	assert.Error(t, err)
}
