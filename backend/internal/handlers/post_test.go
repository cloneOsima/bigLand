package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
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

var (
	lat              = 0.1
	lng              = 0.2
	testUUID, _      = uuid.NewUUID()
	testTime         = time.Now()
	testLocationText = "testlocation"
	sampleText       = "sample text for testing"
)

func TestGetPosts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testCases := []struct {
		name           string
		svcReturn      []*models.Posts
		returnErr      error
		expectedStatus int
		expectedBody   []*models.Posts
	}{
		{
			name: "Success - Formal Response",
			svcReturn: []*models.Posts{
				{PostID: testUUID, PostedDate: testTime, Latitude: &lat, Longtitude: &lng, AddressText: sampleText, Location: []byte(testLocationText)},
				{PostID: testUUID, PostedDate: testTime, Latitude: &lat, Longtitude: &lng, AddressText: sampleText, Location: []byte(testLocationText)},
			},
			returnErr:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: []*models.Posts{
				{PostID: testUUID, PostedDate: testTime, Latitude: &lat, Longtitude: &lng, AddressText: sampleText, Location: []byte(testLocationText)},
				{PostID: testUUID, PostedDate: testTime, Latitude: &lat, Longtitude: &lng, AddressText: sampleText, Location: []byte(testLocationText)},
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
			returnErr:      errors.New("otherErrors"),
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

				// http response 의 body는 []byte 형태의 json 문자열이 들어가 있음.
				// 이를 unmarshal 해서 []models.Posts 형태로 변환한 다음 비교를 진행
				var got []models.Posts
				if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
					t.Fatalf("failed to unmarshal response body: %v", err)
				}

				if len(got) != len(tc.expectedBody) {
					t.Fatalf("expected %d posts, got %d", len(tc.expectedBody), len(got))
				}

				for i := range got {
					expectedLat := *tc.expectedBody[i].Latitude
					expectedLng := *tc.expectedBody[i].Longtitude
					gotLat := *got[i].Latitude
					gotLng := *got[i].Longtitude

					if got[i].PostID != tc.expectedBody[i].PostID {
						t.Errorf("postId mismatch: expected %v, got %v", tc.expectedBody[i].PostID, got[i].PostID)
					}
					if !got[i].PostedDate.Equal(tc.expectedBody[i].PostedDate) {
						t.Errorf("postedDate mismatch: expected %v, got %v", tc.expectedBody[i].PostedDate, got[i].PostedDate)
					}
					if expectedLat != gotLat {
						t.Errorf("latitude mismatch: expected %v, got %v", expectedLat, gotLat)
					}
					if expectedLng != gotLng {
						t.Errorf("longtitude mismatch: expected %v, got %v", expectedLat, gotLat)
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

func TestGetPostInfo(t *testing.T) {

	testUUID, _ := uuid.NewUUID()
	testCases := []struct {
		name           string
		svcReturn      *models.Post
		returnErr      error
		expectedStatus int
		expectedBody   models.Post
	}{
		{
			name: "Success - Formal Result(GetPostInfo)",
			svcReturn: &models.Post{
				PostID:       testUUID,
				Content:      "test-content",
				IncidentDate: testTime,
				PostedDate:   testTime,
				Latitude:     &lat,
				Longtitude:   &lng,
				AddressText:  sampleText,
				Location:     []byte(testLocationText),
				IsActive:     true,
			},
			returnErr:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: models.Post{
				PostID:       testUUID,
				Content:      "test-content",
				IncidentDate: testTime,
				PostedDate:   testTime,
				Latitude:     &lat,
				Longtitude:   &lng,
				AddressText:  sampleText,
				Location:     []byte(testLocationText),
				IsActive:     true,
			},
		},

		{
			name:           "Error - Context RequestTimeout(GetPostInfo)",
			svcReturn:      nil,
			returnErr:      context.DeadlineExceeded,
			expectedStatus: http.StatusRequestTimeout,
			expectedBody:   models.Post{},
		},
		{
			name:           "Error - Context Canceled(GetPostInfo)",
			svcReturn:      nil,
			returnErr:      context.Canceled,
			expectedStatus: http.StatusRequestTimeout,
			expectedBody:   models.Post{},
		},
		{
			name:           "Error - InternalServerError(GetPostInfo)",
			svcReturn:      nil,
			returnErr:      errors.New("otherErrors"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   models.Post{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockSvc := services.NewMockPostService(t)
			mockSvc.On("GetPostInfo", mock.Anything, testUUID.String()).Return(tc.svcReturn, tc.returnErr)
			url := fmt.Sprintf("/post/%s", testUUID)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			req, _ := http.NewRequest(http.MethodGet, url, nil)
			req.Header.Set("X-Request-ID", "test-request-id")

			c.Request = req
			c.Params = gin.Params{
				{Key: "id", Value: testUUID.String()},
			}

			handler := handlers.NewPostHandler(mockSvc)
			handler.GetPostInfo(c)

			if tc.name == "Success - Formal Result(GetPostsInfo)" {
				var got models.Post

				if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
					t.Fatalf("failed to unmarshal response body: %v", err)
				}
				if got.PostID != tc.expectedBody.PostID {
					t.Errorf("postId mismatch: expected %v, got %v", tc.expectedBody.PostID, got.PostID)
				}
				if !got.PostedDate.Equal(tc.expectedBody.PostedDate) {
					t.Errorf("postedDate mismatch: expected %v, got %v", tc.expectedBody.PostedDate, got.PostedDate)
				}
				if got.Latitude != tc.expectedBody.Latitude {
					t.Errorf("latitude mismatch: expected %v, got %v", tc.expectedBody.Latitude, got.Latitude)
				}
				if got.Longtitude != tc.expectedBody.Longtitude {
					t.Errorf("longtitude mismatch: expected %v, got %v", tc.expectedBody.Longtitude, got.Longtitude)
				}
				if got.AddressText != tc.expectedBody.AddressText {
					t.Errorf("addressText mismatch: expected %v, got %v", tc.expectedBody.AddressText, got.AddressText)
				}
				if !bytes.Equal(got.Location, tc.expectedBody.Location) {
					t.Errorf("location mismatch: expected %v, got %v", tc.expectedBody.Location, got.Location)
				}
				if got.Content != tc.expectedBody.Content {
					t.Errorf("content mismatch: expected %v, got %v", tc.expectedBody.Content, got.Content)
				}
				if !got.IncidentDate.Equal(tc.expectedBody.IncidentDate) {
					t.Errorf("incidentDate mismatch: expected %v, got %v", tc.expectedBody.IncidentDate, got.IncidentDate)
				}
				if got.IsActive != tc.expectedBody.IsActive {
					t.Errorf("isActive mismatch: expected %v, got %v", tc.expectedBody.IsActive, got.IsActive)
				}
			}
		})
	}
}
