package services_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/cloneOsima/bigLand/backend/internal/mocks/repositories"
	"github.com/cloneOsima/bigLand/backend/internal/models"
	"github.com/cloneOsima/bigLand/backend/internal/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPosts(t *testing.T) {
	testUUID, _ := uuid.NewUUID()
	testTime := time.Now()
	testLocationText := "testlocation"
	tests := []struct {
		name        string
		mockReturn  []*models.Posts
		mockError   error
		expectPosts []*models.Posts
		expectErr   error
	}{
		{
			name: "Success - Returns posts",
			mockReturn: []*models.Posts{
				{PostId: testUUID, PostedDate: testTime, Latitude: 0.01, Longtitude: 0.02, AddressText: "test", Location: []byte(testLocationText)},
				{PostId: testUUID, PostedDate: testTime.Add(10 * time.Minute), Latitude: 0.02, Longtitude: 0.01, AddressText: "test2", Location: []byte(testLocationText)}},
			mockError: nil,
			expectPosts: []*models.Posts{
				{PostId: testUUID, PostedDate: testTime, Latitude: 0.01, Longtitude: 0.02, AddressText: "test", Location: []byte(testLocationText)},
				{PostId: testUUID, PostedDate: testTime.Add(10 * time.Minute), Latitude: 0.02, Longtitude: 0.01, AddressText: "test2", Location: []byte(testLocationText)}},
			expectErr: nil,
		},
		{
			name:        "Error - failed to query",
			mockReturn:  nil,
			mockError:   assert.AnError,
			expectPosts: nil,
			expectErr:   assert.AnError,
		},
		{
			name:        "Success - No posts found (empty slice)",
			mockReturn:  []*models.Posts{},
			mockError:   nil,
			expectPosts: nil,
			expectErr:   nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockRepo := repositories.NewMockPostRepository(t)
			mockRepo.On("GetPosts", mock.Anything).Return(tc.mockReturn, tc.mockError)

			postService := services.NewPostService(mockRepo)
			result, err := postService.GetPosts(context.Background())

			if (err != nil && tc.expectErr == nil) || (err == nil && tc.expectErr != nil) || (err != nil && err.Error() != tc.expectErr.Error()) {
				t.Errorf("예상 에러: '%v', 실제 에러: '%v'", tc.expectErr, err)
			}
			if len(result) != len(tc.expectPosts) {
				t.Errorf("반환된 게시물 수가 다름. 예상: %d, 실제: %d", len(tc.expectPosts), len(result))
			}
		})
	}
}

func TestGetPostInfo(t *testing.T) {
	testUUID, _ := uuid.NewUUID()
	testTime := time.Now()
	testLocationText := "testlocation"
	tests := []struct {
		name       string
		mockReturn *models.Post
		mockError  error
		expectPost *models.Post
		expectErr  error
	}{
		{
			name:       "Success - Return post info",
			mockReturn: &models.Post{PostId: testUUID, Content: "test-content", IncidentDate: testTime, PostedDate: testTime, Latitude: 0.01, Longtitude: 0.02, AddressText: "test", Location: []byte(testLocationText)},
			mockError:  nil,
			expectPost: &models.Post{PostId: testUUID, Content: "test-content", IncidentDate: testTime, PostedDate: testTime, Latitude: 0.01, Longtitude: 0.02, AddressText: "test", Location: []byte(testLocationText)},
			expectErr:  nil,
		},
		{
			name:       "Error - failed to query",
			mockReturn: nil,
			mockError:  assert.AnError,
			expectPost: &models.Post{},
			expectErr:  assert.AnError,
		},
		{
			name:       "Success - No post found (empty row)",
			mockReturn: nil,
			mockError:  nil,
			expectPost: &models.Post{},
			expectErr:  nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockRepo := repositories.NewMockPostRepository(t)
			mockRepo.On("GetPostInfo", mock.Anything).Return(tc.mockReturn, tc.mockError)

			postService := services.NewPostService(mockRepo)
			result, err := postService.GetPostInfo(context.Background())

			if (err != nil && tc.expectErr == nil) || (err == nil && tc.expectErr != nil) || (err != nil && err.Error() != tc.expectErr.Error()) {
				t.Errorf("예상 에러: '%v', 실제 에러: '%v'", tc.expectErr, err)
			}

			if tc.name == "Success - Return post info" {
				if !reflect.DeepEqual(result, tc.expectPost) {
					t.Errorf("예상 결과: '%v', 실제 결과: '%v'", tc.expectPost, result)
				}
			}
		})
	}
}
