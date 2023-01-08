package queue_test

import (
	"_core/queue"
	"_models/ent"
	"context"
	"fmt"
	"math/rand"
	"testing"

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

func TestCreateQueue(t *testing.T) {
	s, client, l := CreateQueueService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := createUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	queue, err := s.CreateQueue(context.Background(), queue.CreateQueuePayload{
		Name: "some-queue",
		User: usr,
	})
	assert.NotNil(t, queue)
	assert.NoError(t, err)
}
