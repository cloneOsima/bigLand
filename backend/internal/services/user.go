package services

import (
	"context"
	"errors"

	"github.com/cloneOsima/bigLand/backend/internal/models"
	"github.com/cloneOsima/bigLand/backend/internal/repositories"
	"github.com/cloneOsima/bigLand/backend/internal/utils"
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
	err := signUpValidation(inputData)

	// Password hashing

	// mapping models <> sqlc

	// call repo

	// loaded data check

	// return

	return nil
}

func signUpValidation(inputData models.User) error {

	if inputData.Username == "" || len(inputData.Username) <= 5 {
		return errors.New("Empty username")
	}

	if inputData.Email == "" || utils.EmailRegex.MatchString(inputData.Email) {
		return errors.New("Invalid email")
	}

	if inputData.PasswordHash == "" || len(inputData.PasswordHash) <= 5 {
		return errors.New("Invalid password")
	}

	return nil
}
