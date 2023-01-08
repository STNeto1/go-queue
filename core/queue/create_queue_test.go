package queue_test

import (
	"_core/queue"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateQueue(t *testing.T) {
	s, client, l := CreateQueueService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := client.User.
		Create().
		SetName("John Doe").
		SetEmail("some-mail@mail.com").
		SetPassword("some-password").
		Save(context.Background())
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	queue, err := s.CreateQueue(context.Background(), queue.CreateQueuePayload{
		Name: "some-queue",
		User: usr,
	})
	assert.NotNil(t, queue)
	assert.NoError(t, err)
}
