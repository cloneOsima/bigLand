package services_test

import (
	"context"
	"testing"

	"github.com/cloneOsima/bigLand/backend/internal/mocks/repositories"
	"github.com/cloneOsima/bigLand/backend/internal/models"
	"github.com/cloneOsima/bigLand/backend/internal/services"
	"github.com/cloneOsima/bigLand/backend/internal/sqlc"
	"github.com/stretchr/testify/mock"
)

var (
	validInputData = models.User{
		Username:     "testUser",
		Email:        "test@gmail.com",
		PasswordHash: "HashedPass",
	}
	validReturnData = sqlc.SelectUserRow{
		Username:     "testUser",
		Email:        "test@gmail.com",
		PasswordHash: "HashedPass",
	}
)

func TestSignUp(t *testing.T) {
	testCases := []struct {
		name         string
		inputData    models.User
		returnedData sqlc.SelectUserRow
		expectedErr  error
		returnedErr  error
		flag         bool
	}{
		{
			name:         "Success - Create a new account",
			inputData:    validInputData,
			returnedData: validReturnData,
			expectedErr:  nil,
			returnedErr:  nil,
			flag:         true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// repo 설정
			repo := repositories.NewMockUserRepository(t)
			if tc.flag {
				repo.On("InsertNewAccount", mock.Anything, mock.AnythingOfType("sqlc.InsertNewAccountParams")).Return(tc.returnedErr)
			} else {
				repo.AssertNotCalled(t, "InsertNewAccount", mock.Anything, mock.Anything)
			}

			// svc 설정 & 실행
			svc := services.NewUserService(repo)
			ctx := context.Background()
			err := svc.SignUp(ctx, validInputData)

			// 비교
			if tc.expectedErr == nil {
				if err != nil {
					t.Errorf("예상 에러 없음, 실제 에러: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("예상 에러: '%v', 실제 에러: nil", tc.expectedErr)
				} else if err.Error() != tc.expectedErr.Error() {
					t.Errorf("예상 에러: '%v', 실제 에러: '%v'", tc.expectedErr, err)
				}
			}

		})
	}
}
