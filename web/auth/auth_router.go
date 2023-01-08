package auth

import (
	"_core/auth"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	as *auth.AuthService
}

func SetupAuthRoutes(as *auth.AuthService, r *gin.Engine) {
	localRouter := AuthRouter{as: as}

	group := r.Group("/auth")

	group.POST("/login", localRouter.LoginHandler)

}
