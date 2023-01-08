package message_test

import (
	"_core/message"
	"_models/ent"
	"_models/ent/enttest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func CreateMessageService(t *testing.T) (*message.MessageService, *ent.Client, *zap.Logger) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	logger := zap.NewNop()

	return message.NewMessageService(client, logger), client, logger
}

func TestCreateMessageService(t *testing.T) {
	s, client, l := CreateMessageService(t)
	defer client.Close()
	defer l.Sync()

	assert.NotNil(t, s)
	assert.NotNil(t, client)
}
