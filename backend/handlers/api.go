package handlers

// Handlers package for processing Http request.
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cloneOsima/bigLand/backend/services"
	"github.com/cloneOsima/bigLand/backend/utils"
	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetPostList(c *gin.Context) {

	reqId := c.GetHeader("X-Request-ID")
	if reqId == "" {
		reqId = fmt.Sprint(utils.GenerateRequestId())
	}

	// use gin.Contexts -> can check user connection & middleware information
	ctx, cancle := context.WithTimeout(c.Request.Context(), 10*time.Second)
	ctx = context.WithValue(ctx, "requestId", reqId)
	defer cancle()

	result, err := services.GetPostList(ctx)
	if err != nil {
		log.Printf("GetPostList failed - Request Id: %s, Error, %v", reqId, err)

		if ctx.Err() == context.DeadlineExceeded {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"error":     "Request timeout",
				"requestID": reqId,
			})
			return
		}

		if ctx.Err() == context.Canceled {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"error":     "Request cancelled",
				"requestID": reqId,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":     err.Error(),
			"requestID": reqId,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
