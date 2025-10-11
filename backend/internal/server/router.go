package server

// Server package for initializing Gin router and server

import (
	"github.com/cloneOsima/bigLand/backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

// Ping: 각 handler 정상 연결 확인용 health checker
func SetupRouter(h *handlers.Handlers) *gin.Engine {
	r := gin.Default()

	// Post package routing
	r.GET("/post/ping", h.Post.Ping)
	r.GET("/posts", h.Post.GetPosts)
	r.GET("/post/:id", h.Post.GetPostInfo)
	r.POST("/post", h.Post.CreatePost)

	// User pacakge routing
	r.GET("/user/ping", h.User.Ping)
	r.POST("/user", h.User.CreateAccount)

	return r
}
