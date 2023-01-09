package web

import (
	"_core/auth"
	lib "_lib"
	ar "_web/auth"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Router struct {
	Engine *echo.Echo

	AS *auth.AuthService
}

func NewRouter(AS *auth.AuthService) *Router {
	ec := echo.New()
	ec.Validator = &lib.CustomValidator{Validator: validator.New()}

	ar.SetupAuthRoutes(AS, ec)

	return &Router{
		Engine: ec,
		AS:     AS,
	}

}
