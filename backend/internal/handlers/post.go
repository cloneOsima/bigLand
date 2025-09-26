package handlers

// Handlers package for processing Http request.
import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	errdefs "github.com/cloneOsima/bigLand/backend/internal/errors"
	"github.com/cloneOsima/bigLand/backend/internal/models"
	"github.com/cloneOsima/bigLand/backend/internal/services"
	"github.com/cloneOsima/bigLand/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetPosts(c *gin.Context)
	GetPostInfo(c *gin.Context)
	CreatePost(c *gin.Context)
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

// 신규 게시글 투고
func (h *handlerImpl) CreatePost(c *gin.Context) {
	reqId := reqIdChecker(c)
	var inputValue models.Post

	var requestIdKey utils.CtxKey = "reqId"
	ctx, cancle := context.WithTimeout(c.Request.Context(), 10*time.Second)
	ctx = context.WithValue(ctx, requestIdKey, reqId)
	defer cancle()

	// request body <> models package mapping
	if err := c.ShouldBindJSON(&inputValue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.postSvc.CreatePost(ctx, &inputValue)
	commonErrorChecker(c, ctx, err, reqId)
	// 성공 시 200 status code + 입력값 반환(프론트 구성시 필요 없으면 nil 등으로 처리 할 것)
	c.JSON(http.StatusOK, inputValue)
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

		// 에러가 잘못된 값으로 발생한 에러일 경우 400과 함께 내용 반환
		if apperr, ok := err.(*errdefs.ValueErr); ok {
			log.Printf("[Request failed] Request Id: %s, Error: %v, ErrorInfo: %v", reqId, err, apperr.ErrorInfo)
			c.JSON(http.StatusBadRequest, gin.H{
				"error":     "Request timeout",
				"requestID": reqId,
			})
			return
		}

		log.Printf("[Request failed] Request Id: %s, Error: %v", reqId, err)
		if errors.Is(err, context.DeadlineExceeded) || ctx.Err() == context.DeadlineExceeded {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"error":     "Request timeout",
				"requestID": reqId,
			})
			return
		}
		if errors.Is(err, context.Canceled) || ctx.Err() == context.Canceled {
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
