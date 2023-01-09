package auth

import (
	"_core/auth"

	"github.com/labstack/echo/v4"
)

type AuthRouter struct {
	as *auth.AuthService
}

func SetupAuthRoutes(as *auth.AuthService, r *echo.Echo) {
	localRouter := AuthRouter{as: as}

	group := r.Group("/auth")

	group.POST("/login", localRouter.LoginHandler)
	group.POST("/register", localRouter.RegisterHandler)

}
