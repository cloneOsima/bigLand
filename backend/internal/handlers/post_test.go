package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/cloneOsima/bigLand/backend/internal/handlers"
	"github.com/cloneOsima/bigLand/backend/internal/mocks/services"
	"github.com/cloneOsima/bigLand/backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func TestGetPosts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testUUID, _ := uuid.NewUUID()
	testTime := time.Now()
	testLocationText := "testlocation"

	testCases := []struct {
		name           string
		svcReturn      []models.Posts
		withCtxFunc    func(ctx context.Context) context.Context
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success - Formal response",
			svcReturn: []models.Posts{
				{PostId: testUUID, PostedDate: testTime, Latitude: 0.01, Longtitude: 0.02, AddressText: "test", Location: []byte(testLocationText)},
				{PostId: testUUID, PostedDate: testTime.Add(10 * time.Minute), Latitude: 0.02, Longtitude: 0.01, AddressText: "test2", Location: []byte(testLocationText)},
			},
			withCtxFunc:    nil,
			expectedStatus: http.StatusOK,
			expectedBody: `[
				{
					"post_id":"` + testUUID.String() + `",
					"posted_date":"` + testTime.Format(time.RFC3339Nano) + `",
					"latitude":0.01,
					"longtitude":0.02,
					"address_text":"test",
					"location":"dGVzdGxvY2F0aW9u"
				},
				{
					"post_id":"` + testUUID.String() + `",
					"posted_date":"` + testTime.Add(10*time.Minute).Format(time.RFC3339Nano) + `",
					"latitude":0.02,
					"longtitude":0.01,
					"address_text":"test2",
					"location":"dGVzdGxvY2F0aW9u"
				}
			]`,
		},
		{
			name:      "Error - ContextDeadlineExceeded",
			svcReturn: nil,
			withCtxFunc: func(ctx context.Context) context.Context {
				ctx, cancel := context.WithTimeout(ctx, 1*time.Millisecond)
				time.Sleep(2 * time.Millisecond)
				cancel()
				return ctx
			},
			expectedStatus: http.StatusRequestTimeout,
			expectedBody:   `{"error":"Request timeout","requestID":"test-request-id"}`,
		},
		{
			name:      "Error - ContextCanceled",
			svcReturn: nil,
			withCtxFunc: func(ctx context.Context) context.Context {
				ctx, cancel := context.WithCancel(ctx)
				cancel()
				return ctx
			},
			expectedStatus: http.StatusRequestTimeout,
			expectedBody:   `{"error":"Request cancelled","requestID":"test-request-id"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockSvc := services.NewMockPostService(t)
			mockSvc.On("GetPosts", mock.Anything).Return(tc.svcReturn, nil)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			req, _ := http.NewRequest(http.MethodGet, "/posts", nil)
			req.Header.Set("X-Request-ID", "test-request-id")

			// Context 적용
			if tc.withCtxFunc != nil {
				req = req.WithContext(tc.withCtxFunc(req.Context()))
			}
			c.Request = req

			handler := handlers.NewPostHandler(mockSvc)
			handler.GetPosts(c)

			// 상태 코드 검증
			if w.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Code)
			}

			// JSON 비교
			if tc.name == "Success - Formal response" {
				var got, expected []interface{}
				if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
					t.Fatalf("failed to unmarshal response body: %v", err)
				}
				if err := json.Unmarshal([]byte(tc.expectedBody), &expected); err != nil {
					t.Fatalf("failed to unmarshal expected body: %v", err)
				}
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("expected body %v, got %v", expected, got)
				}
			} else {
				var got, expected map[string]interface{}
				if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
					t.Fatalf("failed to unmarshal response body: %v", err)
				}
				if err := json.Unmarshal([]byte(tc.expectedBody), &expected); err != nil {
					t.Fatalf("failed to unmarshal expected body: %v", err)
				}
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("expected body %v, got %v", expected, got)
				}
			}
		})
	}
}
