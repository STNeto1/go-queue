package web_test

import (
	"_core/auth"
	"_models/ent/enttest"
	web "_web"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestSetupRouter(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	logger := zap.NewNop()
	defer client.Close()
	defer logger.Sync()

	as := auth.NewAuthService(client, "some-secret", logger)
	assert.NotNil(t, as)

	router := web.NewRouter(as, logger)

	assert.NotNil(t, router)
}
