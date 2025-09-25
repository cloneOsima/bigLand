package handlers

// Handlers package for processing Http request.
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cloneOsima/bigLand/backend/internal/services"
	"github.com/cloneOsima/bigLand/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetPosts(c *gin.Context)
	GetPostInfo(c *gin.Context)
}

type handlerImpl struct {
	postSvc services.PostService
}

func NewPostHandler(postSvc services.PostService) Handler {
	return &handlerImpl{
		postSvc: postSvc,
	}
}

func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// 등록일 기준 최근 50개를 호출하는 함수 (페이지 로딩시 자동 표시)
func (h *handlerImpl) GetPosts(c *gin.Context) {

	reqId := reqIdChecker(c)
	var requestIdKey utils.CtxKey = "reqId"
	// use gin.Contexts -> can check user connection & middleware information
	ctx, cancle := context.WithTimeout(c.Request.Context(), 10*time.Second)
	ctx = context.WithValue(ctx, requestIdKey, reqId)
	defer cancle()

	result, err := h.postSvc.GetPosts(ctx)
	commonErrorChecker(c, ctx, err, reqId)
	c.JSON(http.StatusOK, result)
}

// 특정 GIS 데이터를 선택하거나, 최근 등록글을 선택하면 해당 투고에 대한 상세 정보를 제공해주는 함수
func (h *handlerImpl) GetPostInfo(c *gin.Context) {
	reqId := reqIdChecker(c)
	postID := c.Param("id")

	var requestIdKey utils.CtxKey = "reqId"
	ctx, cancle := context.WithTimeout(c.Request.Context(), 10*time.Second)
	ctx = context.WithValue(ctx, requestIdKey, reqId)
	defer cancle()

	result, err := h.postSvc.GetPostInfo(ctx, postID)
	commonErrorChecker(c, ctx, err, reqId)
	c.JSON(http.StatusOK, result)
}

// request id 관련 로직을 통합해서 처리해주는 함수
func reqIdChecker(c *gin.Context) string {
	reqId := c.GetHeader("X-Request-ID")
	if reqId == "" {
		reqId = fmt.Sprint(utils.GenerateRequestId())
	}

	return reqId
}

// 공통적으로 발생할 수 있는 http response 관련 에러를 처리하는 함수
func commonErrorChecker(c *gin.Context, ctx context.Context, err error, reqId string) {
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

}
