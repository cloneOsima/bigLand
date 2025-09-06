package handlers

// Handlers package for processing Http request.
import (
	"net/http"

	"github.com/cloneOsima/bigLand/backend/services"
	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetPostList(c *gin.Context) {
	result, err := services.GetPostList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, result)
}
