package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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
	testLocationText := []byte("testlocation")
	otherErrors := fmt.Errorf("Someting goes wrong except ctx.Deadline and ctx.Canceled.")

	testCases := []struct {
		name           string
		svcReturn      []models.Posts
		returnErr      error
		expectedStatus int
		expectedBody   []models.Posts
	}{
		{
			name: "Success - Formal Response",
			svcReturn: []models.Posts{
				{PostId: testUUID, PostedDate: testTime, Latitude: 0.01, Longtitude: 0.02, AddressText: "test", Location: testLocationText},
				{PostId: testUUID, PostedDate: testTime, Latitude: 0.02, Longtitude: 0.01, AddressText: "test2", Location: testLocationText},
			},
			returnErr:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: []models.Posts{
				{PostId: testUUID, PostedDate: testTime, Latitude: 0.01, Longtitude: 0.02, AddressText: "test", Location: testLocationText},
				{PostId: testUUID, PostedDate: testTime, Latitude: 0.02, Longtitude: 0.01, AddressText: "test2", Location: testLocationText},
			},
		},
		{
			name:           "Error - Context RequestTimeout",
			svcReturn:      nil,
			returnErr:      context.DeadlineExceeded,
			expectedStatus: http.StatusRequestTimeout,
			expectedBody:   nil,
		},
		{
			name:           "Error - Context Canceled",
			svcReturn:      nil,
			returnErr:      context.Canceled,
			expectedStatus: http.StatusRequestTimeout,
			expectedBody:   nil,
		},
		{
			name:           "Error - InternalServerError",
			svcReturn:      nil,
			returnErr:      otherErrors,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockSvc := services.NewMockPostService(t)
			mockSvc.On("GetPosts", mock.Anything).Return(tc.svcReturn, tc.returnErr)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			req, _ := http.NewRequest(http.MethodGet, "/posts", nil)
			req.Header.Set("X-Request-ID", "test-request-id")

			c.Request = req

			handler := handlers.NewPostHandler(mockSvc)
			handler.GetPosts(c)

			if tc.name == "Success - Formal Response" {
				var got []models.Posts
				if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
					t.Fatalf("failed to unmarshal response body: %v", err)
				}

				if len(got) != len(tc.expectedBody) {
					t.Fatalf("expected %d posts, got %d", len(tc.expectedBody), len(got))
				}

				for i := range got {
					if got[i].PostId != tc.expectedBody[i].PostId {
						t.Errorf("postId mismatch: expected %v, got %v", tc.expectedBody[i].PostId, got[i].PostId)
					}
					if !got[i].PostedDate.Equal(tc.expectedBody[i].PostedDate) {
						t.Errorf("postedDate mismatch: expected %v, got %v", tc.expectedBody[i].PostedDate, got[i].PostedDate)
					}
					if got[i].Latitude != tc.expectedBody[i].Latitude {
						t.Errorf("latitude mismatch: expected %v, got %v", tc.expectedBody[i].Latitude, got[i].Latitude)
					}
					if got[i].Longtitude != tc.expectedBody[i].Longtitude {
						t.Errorf("longtitude mismatch: expected %v, got %v", tc.expectedBody[i].Longtitude, got[i].Longtitude)
					}
					if got[i].AddressText != tc.expectedBody[i].AddressText {
						t.Errorf("addressText mismatch: expected %v, got %v", tc.expectedBody[i].AddressText, got[i].AddressText)
					}
					if !bytes.Equal(got[i].Location, tc.expectedBody[i].Location) {
						t.Errorf("location mismatch: expected %v, got %v", tc.expectedBody[i].Location, got[i].Location)
					}
				}
			}
		})
	}
}
