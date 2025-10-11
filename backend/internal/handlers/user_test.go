package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloneOsima/bigLand/backend/internal/handlers"
	"github.com/cloneOsima/bigLand/backend/internal/mocks/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

var (
	validInput = `{"username": "testUser", "email": "test@gmail.com", "password": "testPassword1!@"}`
)

func TestCreateAccount(t *testing.T) {
	testCases := []struct {
		name           string
		expectedErr    error
		returnErr      error
		expectedStatus int
	}{
		{
			name:           "Success - create new account",
			expectedErr:    nil,
			returnErr:      nil,
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// svc mock 객체 생성
			svc := services.NewMockUserService(t)
			svc.On("SignUp", mock.Anything, mock.AnythingOfType("models.User")).Return(tc.returnErr)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer([]byte(validInput)))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-Request-ID", "test-request-id")
			c.Request = req

			handler := handlers.NewUserHandler(svc)
			handler.CreateAccount(c)

			if w.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Code)
			}
		})
	}
}
