// server 초기화 시 service 를 한번에 묶어서 초기화
package services

import "github.com/cloneOsima/bigLand/backend/internal/repositories"

type Services struct {
	User UserService
	Post PostService
}

func InitSvc(repo *repositories.Repositories) *Services {
	return &Services{
		User: NewUserService(repo.User),
		Post: NewPostService(repo.Post),
	}
}
