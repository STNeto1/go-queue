package web

import (
	"_core/auth"
	ar "_web/auth"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine

	AS *auth.AuthService
}

func NewRouter(AS *auth.AuthService) *Router {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	ar.SetupAuthRoutes(AS, r)

	return &Router{
		Engine: r,
		AS:     AS,
	}

}
