package handlers

// Handlers package for processing Http request.
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
