package auth_test

import (
	"_core/auth"
	"_models/ent"
	"_models/ent/enttest"
	web "_web"
	am "_web/auth"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func createAuthRouter(t *testing.T) (*web.Router, *ent.Client) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	logger := zap.NewNop()

	as := auth.NewAuthService(client, "some-secret", logger)
	assert.NotNil(t, as)

	router := web.NewRouter(as, logger)

	assert.NotNil(t, router)

	return router, client
}

func createUser(t *testing.T, client *ent.Client) *ent.User {
	existingMail := fmt.Sprintf("some-%d@mail.com", rand.Int())
	correctPassword := "some-password"

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(correctPassword), bcrypt.DefaultCost)
	assert.NoError(t, err)

	usr, err := client.User.Create().SetName("John Doe").SetEmail(existingMail).SetPassword(string(pwdHash)).Save(context.Background())
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	return usr
}

func TestLoginHandlerBadValidation(t *testing.T) {
	router, client := createAuthRouter(t)
	defer client.Close()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/auth/login", nil)
	router.Engine.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestLoginHandlerBadCredentials(t *testing.T) {
	router, client := createAuthRouter(t)
	defer client.Close()

	w := httptest.NewRecorder()

	body := am.LoginRequestBody{
		Email:    "some-mail@mail.com",
		Password: "some password",
	}
	jsonValue, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonValue))

	router.Engine.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestLoginHandlerSuccess(t *testing.T) {
	router, client := createAuthRouter(t)
	defer client.Close()

	usr := createUser(t, client)

	body := am.LoginRequestBody{
		Email:    usr.Email,
		Password: "some-password",
	}
	jsonVal, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(string(jsonVal)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	router.Engine.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
}
