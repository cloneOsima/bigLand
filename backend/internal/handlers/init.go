// server 초기화 시 handler 를 한번에 묶어서 초기화
package handlers

import "github.com/cloneOsima/bigLand/backend/internal/services"

type Handlers struct {
	User UserHandler
	Post PostHandler
}

func InitHandler(svc *services.Services) *Handlers {
	return &Handlers{
		User: NewUserHandler(svc.User),
		Post: NewPostHandler(svc.Post),
	}
}
