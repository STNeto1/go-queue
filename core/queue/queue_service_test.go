package queue_test

import (
	"_core/queue"
	"_models/ent"
	"_models/ent/enttest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func CreateQueueService(t *testing.T) (*queue.QueueService, *ent.Client, *zap.Logger) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	logger := zap.NewNop()

	return queue.NewQueueService(client, logger), client, logger
}

func TestCreateQueueService(t *testing.T) {
	s, client, l := CreateQueueService(t)
	defer client.Close()
	defer l.Sync()

	assert.NotNil(t, s)
	assert.NotNil(t, client)
}
