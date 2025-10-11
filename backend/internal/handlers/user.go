package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/cloneOsima/bigLand/backend/internal/models"
	"github.com/cloneOsima/bigLand/backend/internal/services"
	"github.com/cloneOsima/bigLand/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateAccount(c *gin.Context)
	Ping(c *gin.Context)
}

type userHandlerImpl struct {
	svc services.UserService
}

func NewUserHandler(userSvc services.UserService) UserHandler {
	return &userHandlerImpl{
		svc: userSvc,
	}
}

func (u *userHandlerImpl) Ping(c *gin.Context) {
	c.String(http.StatusOK, "user hadler - pong")
}

func (u *userHandlerImpl) CreateAccount(c *gin.Context) {
	reqId := reqIdChecker(c)
	var requestIdKey utils.CtxKey = "reqId"
	var inputValue models.User

	ctx, cancle := context.WithTimeout(c.Request.Context(), 10*time.Second)
	ctx = context.WithValue(ctx, requestIdKey, reqId)
	defer cancle()

	if err := c.ShouldBindJSON(&inputValue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := u.svc.SignUp(ctx, inputValue)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, inputValue)
}
