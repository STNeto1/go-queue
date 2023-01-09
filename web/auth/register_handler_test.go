package auth_test

import (
	am "_web/auth"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandlerBadValidation(t *testing.T) {
	router, client := createAuthRouter(t)
	defer client.Close()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/auth/register", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	router.Engine.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestRegisterHandlerExistingUser(t *testing.T) {
	router, client := createAuthRouter(t)
	defer client.Close()

	usr := createUser(t, client)

	body := am.RegisterRequestBody{
		Name:     "John Doe",
		Email:    usr.Email,
		Password: "some-password",
	}
	jsonVal, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(string(jsonVal)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	router.Engine.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestRegisterHandlerSuccess(t *testing.T) {
	router, client := createAuthRouter(t)
	defer client.Close()

	body := am.RegisterRequestBody{
		Name:     "John Doe",
		Email:    "some-mail@mail.com",
		Password: "some-password",
	}
	jsonVal, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(string(jsonVal)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	router.Engine.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
}
