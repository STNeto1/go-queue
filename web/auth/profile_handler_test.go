package auth_test

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProfileHandlerNoToken(t *testing.T) {
	router, client := createAuthRouter(t)
	defer client.Close()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/auth/profile", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	router.Engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestProfileHandlerSuccess(t *testing.T) {
	router, client := createAuthRouter(t)
	defer client.Close()

	usr := createUser(t, client)

	token, _ := router.AS.GenerateToken(usr)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/auth/profile", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	router.Engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
