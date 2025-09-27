package handlers

import (
	"net/http"

	"github.com/cloneOsima/bigLand/backend/internal/services"
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
	// reqId := reqIdChecker(c)
	// var requestIdKey utils.CtxKey = "reqId"
	// // use gin.Contexts -> can check user connection & middleware information
	// ctx, cancle := context.WithTimeout(c.Request.Context(), 10*time.Second)
	// ctx = context.WithValue(ctx, requestIdKey, reqId)
	// defer cancle()

	// c.JSON(http.StatusOK, "ping-pong-user")
}
