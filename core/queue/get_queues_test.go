package queue_test

import (
	qm "_core/queue"
	"_models/ent"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createQueue(client *ent.Client, usr *ent.User, name string) (*ent.Queue, error) {
	return client.Queue.
		Create().
		SetName(name).
		SetUser(usr).
		Save(context.Background())
}

func TestGetQueuesWithNoName(t *testing.T) {
	s, client, l := CreateQueueService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	queue, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, queue)
	assert.NoError(t, err)

	queues, err := s.GetQueues(context.Background(), qm.GetQueuesPayload{
		Name: "",
		User: usr,
	})

	assert.NotNil(t, queues)
	assert.Len(t, queues, 1)
	assert.NoError(t, err)
}

func TestGetQueuesWithName(t *testing.T) {
	s, client, l := CreateQueueService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	queue, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, queue)
	assert.NoError(t, err)

	queues, err := s.GetQueues(context.Background(), qm.GetQueuesPayload{
		Name: "complex queue name",
		User: usr,
	})

	assert.NotNil(t, queues)
	assert.Len(t, queues, 0)
	assert.NoError(t, err)
}
