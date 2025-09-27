package handlers

// Handlers package for processing Http request.
import (
	"context"
	"net/http"
	"time"

	"github.com/cloneOsima/bigLand/backend/internal/models"
	"github.com/cloneOsima/bigLand/backend/internal/services"
	"github.com/cloneOsima/bigLand/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type PostHandler interface {
	Ping(c *gin.Context)
	GetPosts(c *gin.Context)
	GetPostInfo(c *gin.Context)
	CreatePost(c *gin.Context)
}

type postHandlerImpl struct {
	postSvc services.PostService
}

func NewPostHandler(postSvc services.PostService) PostHandler {
	return &postHandlerImpl{
		postSvc: postSvc,
	}
}

func (p *postHandlerImpl) Ping(c *gin.Context) {
	c.String(http.StatusOK, "post handler - pong")
}

// 등록일 기준 최근 50개를 호출하는 함수 (페이지 로딩시 자동 표시)
func (h *postHandlerImpl) GetPosts(c *gin.Context) {

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
func (h *postHandlerImpl) GetPostInfo(c *gin.Context) {
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
func (h *postHandlerImpl) CreatePost(c *gin.Context) {
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

	err := h.postSvc.NewPost(ctx, &inputValue)
	commonErrorChecker(c, ctx, err, reqId)
	// 성공 시 200 status code + 입력값 반환(프론트 구성시 필요 없으면 nil 등으로 처리 할 것)
	c.JSON(http.StatusOK, inputValue)
}
