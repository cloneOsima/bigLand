package server

// Server package for initializing Gin router and server

import (
	"github.com/cloneOsima/bigLand/backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", handlers.PingHandler)

	// Routes for Post functions
	r.GET("/posts", handlers.GetPostList)

	return r
}
