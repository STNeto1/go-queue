package web

import "github.com/gin-gonic/gin"

type Router struct {
	engine *gin.Engine
}

func NewRouter() *Router {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return &Router{
		engine: r,
	}
}
