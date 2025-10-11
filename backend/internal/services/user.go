package services

import (
	"context"
	"errors"

	"github.com/cloneOsima/bigLand/backend/internal/models"
	"github.com/cloneOsima/bigLand/backend/internal/repositories"
	"github.com/cloneOsima/bigLand/backend/internal/sqlc"
	"github.com/cloneOsima/bigLand/backend/internal/utils"
	"golang.org/x/crypto/bcrypt"
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
	if err != nil {
		return err
	}

	// Password hashing
	hashPass, err := bcrypt.GenerateFromPassword([]byte(inputData.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// mapping models <> sqlc
	sqlcValue := sqlc.InsertNewAccountParams{}
	sqlcValue.Username = inputData.Username
	sqlcValue.Email = inputData.Email
	sqlcValue.PasswordHash = string(hashPass)

	// db context
	dbCtx, cancel := context.WithTimeout(ctx, defaultDBTimeout)
	defer cancel()

	// call repo
	err = u.userRepo.InsertNewAccount(dbCtx, sqlcValue)
	if err != nil {
		return err
	}

	return nil
}

func signUpValidation(inputData models.User) error {

	if inputData.Username == "" {
		return errors.New("empty username")
	}

	if inputData.Email == "" {
		return errors.New("empty email")
	}

	if !utils.EmailRegex.MatchString(inputData.Email) {
		return errors.New("invalid email")
	}

	err := utils.ValidatePassword(inputData.Password)
	if err != nil {
		return err
	}

	return nil
}
