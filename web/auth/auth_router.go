package auth

import (
	"_core/auth"
	"_models/ent"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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
	group.GET("/profile", localRouter.ProfileHandler)
}

func Skipper(c echo.Context) bool {
	if c.Path() == "/auth/login" || c.Path() == "/auth/register" {
		return true
	}

	return false
}

func (router AuthRouter) UserFromContext(c echo.Context) (*ent.User, error) {
	ctxVal := c.Get("user")
	if ctxVal == nil {
		return nil, errors.New("no value in context")
	}

	claims, ok := ctxVal.(*jwt.StandardClaims)
	if !ok {
		return nil, errors.New("error while parsing claims")
	}

	_id, _ := uuid.Parse(claims.Subject)
	usr, err := router.as.GetUserFromId(c.Request().Context(), _id)
	if err != nil {
		return nil, errors.New("error while getting user from id")
	}

	return usr, nil
}
