package queue_test

import (
	qm "_core/queue"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteQueue(t *testing.T) {
	s, client, l := CreateQueueService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	queue, err := createQueue(client, usr, "some-queue")
	assert.NotNil(t, queue)
	assert.NoError(t, err)

	err = s.DeleteQueue(context.Background(), qm.DeleteQueuePayload{
		ID:   queue.ID,
		User: usr,
	})
	assert.NoError(t, err)
}
