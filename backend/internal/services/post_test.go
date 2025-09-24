package services_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/cloneOsima/bigLand/backend/internal/mocks/repositories"
	"github.com/cloneOsima/bigLand/backend/internal/models"
	"github.com/cloneOsima/bigLand/backend/internal/services"
	"github.com/cloneOsima/bigLand/backend/internal/sqlc"
	"github.com/cloneOsima/bigLand/backend/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPosts(t *testing.T) {
	testUUID, _ := uuid.NewUUID()
	testTime := time.Now()
	testLocationText := "testlocation"
	tests := []struct {
		name        string
		mockReturn  []sqlc.GetPostsRow
		mockError   error
		expectPosts []*models.Posts
		expectErr   error
	}{
		{
			name: "Success - Returns posts",
			mockReturn: []sqlc.GetPostsRow{
				{PostID: testUUID, PostedDate: testTime, Latitude: 0.01, Longtitude: 0.02, AddressText: "test", Location: []byte(testLocationText)},
				{PostID: testUUID, PostedDate: testTime.Add(10 * time.Minute), Latitude: 0.02, Longtitude: 0.01, AddressText: "test2", Location: []byte(testLocationText)}},
			mockError: nil,
			expectPosts: []*models.Posts{
				{PostID: testUUID, PostedDate: testTime, Latitude: 0.01, Longtitude: 0.02, AddressText: "test", Location: []byte(testLocationText)},
				{PostID: testUUID, PostedDate: testTime.Add(10 * time.Minute), Latitude: 0.02, Longtitude: 0.01, AddressText: "test2", Location: []byte(testLocationText)}},
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
			mockReturn:  []sqlc.GetPostsRow{},
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
	testDate := time.Date(2025, 9, 24, 0, 0, 0, 0, time.UTC)
	pgDate := pgtype.Date{
		Time:  testDate,
		Valid: true,
	}
	testLocationText := "testlocation"
	var testCtxKey utils.CtxKey = "postId"

	tests := []struct {
		name       string
		mockReturn sqlc.GetPostInfoRow
		mockError  error
		expectPost *models.Post
		expectErr  error
	}{
		{
			name:       "Success - Return post info",
			mockReturn: sqlc.GetPostInfoRow{PostID: testUUID, Content: "test-content", IncidentDate: pgDate, PostedDate: testTime, Latitude: 0.01, Longtitude: 0.02, AddressText: "test", Location: []byte(testLocationText)},
			mockError:  nil,
			expectPost: &models.Post{PostID: testUUID, Content: "test-content", IncidentDate: testDate, PostedDate: testTime, Latitude: 0.01, Longtitude: 0.02, AddressText: "test", Location: []byte(testLocationText)},
			expectErr:  nil,
		},
		{
			name:       "Error - failed to query",
			mockReturn: sqlc.GetPostInfoRow{},
			mockError:  assert.AnError,
			expectPost: &models.Post{},
			expectErr:  assert.AnError,
		},
		{
			name:       "Success - No post found (empty row)",
			mockReturn: sqlc.GetPostInfoRow{},
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

			ctx := context.Background()
			ctx = context.WithValue(ctx, testCtxKey, testUUID.String())

			postService := services.NewPostService(mockRepo)
			result, err := postService.GetPostInfo(ctx)

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
