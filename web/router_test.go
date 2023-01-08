package web_test

import (
	web "_web"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
	router := web.NewRouter()

	assert.NotNil(t, router)
}
