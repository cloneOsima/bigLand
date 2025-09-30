package services

import (
	"context"

	"github.com/cloneOsima/bigLand/backend/internal/models"
	"github.com/cloneOsima/bigLand/backend/internal/repositories"
)

type UserService interface {
	SignUp(ctx context.Context, data models.User) error
}

type userServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userServiceImpl{userRepo: repo}
}

func (u *userServiceImpl) SignUp(ctx context.Context, inputData models.User) error {

	// input data valid check

	// mapping models <> sqlc

	// call repo

	// loaded data check

	// return

	return nil
}
