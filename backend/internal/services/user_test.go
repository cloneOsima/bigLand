package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/cloneOsima/bigLand/backend/internal/mocks/repositories"
	"github.com/cloneOsima/bigLand/backend/internal/models"
	"github.com/cloneOsima/bigLand/backend/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	inputData = models.User{
		Username: "testUser",
		Email:    "test@gmail.com",
		Password: "testPassword1!@",
	}
)

// 회원 가입 테스트 케이스
// 1. 성공 - 정상 가입
// 2. 실패 - 입력값 검증 실패 (repo 호출 x)
// 3. 실패 - context timeout
func TestSignUp(t *testing.T) {

	testCases := []struct {
		name        string
		inputData   models.User
		expectedErr error
		returnedErr error
		flag        bool
	}{
		{
			name:        "Success - Create a new account",
			inputData:   inputData,
			expectedErr: nil,
			returnedErr: nil,
			flag:        true,
		},
		{
			name:        "Error - Invalid username (empty value)",
			inputData:   models.User{Username: "", Email: inputData.Email, Password: inputData.Password},
			expectedErr: errors.New("empty username"),
			returnedErr: errors.New("empty username"),
			flag:        false,
		},
		{
			name:        "Error - Invalid email (empty value)",
			inputData:   models.User{Username: inputData.Username, Email: "", Password: inputData.Password},
			expectedErr: errors.New("empty email"),
			returnedErr: errors.New("empty email"),
			flag:        false,
		},
		{
			name:        "Error - Invalid email (invalid value)",
			inputData:   models.User{Username: inputData.Username, Email: "test@test", Password: inputData.Password},
			expectedErr: errors.New("invalid email"),
			returnedErr: errors.New("invalid email"),
			flag:        false,
		},
		{
			name:        "Error - Context timeout",
			inputData:   models.User{Username: inputData.Username, Email: "test@test", Password: inputData.Password},
			expectedErr: assert.AnError,
			returnedErr: assert.AnError,
			flag:        false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// repo 설정
			repo := repositories.NewMockUserRepository(t)

			if tc.flag {
				repo.On("InsertNewAccount", mock.Anything, mock.AnythingOfType("sqlc.InsertNewAccountParams")).Return(tc.returnedErr)
			}

			// svc 설정 & 실행
			svc := services.NewUserService(repo)
			ctx := context.Background()
			err := svc.SignUp(ctx, tc.inputData)

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

			if !tc.flag {

				// AssertNotCalled는 mock 호출이 종료후, 해당 method의 호출 여부에 따라 테스트의 성공/실패를 확인함
				repo.AssertNotCalled(t, "InsertNewAccount", mock.Anything, mock.Anything)
			} else {
				repo.AssertCalled(t, "InsertNewAccount", mock.Anything, mock.AnythingOfType("sqlc.InsertNewAccountParams"))
			}
		})
	}
}
