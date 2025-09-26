package server

// Server package for initializing Gin router and server

import (
	"github.com/cloneOsima/bigLand/backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h handlers.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handlers.PingHandler)
	r.GET("/posts", h.GetPosts)
	r.GET("/post/:id", h.GetPostInfo)
	r.POST("/post", h.CreatePost)

	return r
}
