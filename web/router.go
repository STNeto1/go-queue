package web

import (
	"_core/auth"
	lib "_lib"
	ar "_web/auth"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Router struct {
	Engine *echo.Echo

	AS *auth.AuthService
}

func NewRouter(AS *auth.AuthService, logger *zap.Logger) *Router {
	ec := echo.New()
	ec.Validator = &lib.CustomValidator{Validator: validator.New()}

	ec.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "Authorization"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	ec.Use(middleware.Recover())
	ec.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(AS.GetSecret()),
		TokenLookup: "header:Authorization",
		//ContinueOnIgnoredError: false,
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			tokens := strings.Split(auth, " ")
			if len(tokens) != 2 {
				return nil, errors.New("invalid token")
			}

			return AS.ValidateToken(tokens[1])
		},
		Skipper: ar.Skipper,
	}))
	ec.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	ar.SetupAuthRoutes(AS, ec)

	return &Router{
		Engine: ec,
		AS:     AS,
	}

}
