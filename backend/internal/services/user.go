package services

import (
	"context"

	"github.com/cloneOsima/bigLand/backend/internal/repositories"
)

type UserService interface {
	SignUp(ctx context.Context) error
}

type userServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userServiceImpl{userRepo: repo}
}

func (u *userServiceImpl) SignUp(ctx context.Context) error {
	
	return nil
}
