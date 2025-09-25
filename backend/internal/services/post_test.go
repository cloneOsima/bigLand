package services_test

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	errdefs "github.com/cloneOsima/bigLand/backend/internal/errors"
	"github.com/cloneOsima/bigLand/backend/internal/mocks/repositories"
	"github.com/cloneOsima/bigLand/backend/internal/models"
	"github.com/cloneOsima/bigLand/backend/internal/services"
	"github.com/cloneOsima/bigLand/backend/internal/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	lat              = 0.1
	lng              = 0.2
	testUUID, _      = uuid.NewUUID()
	testTime         = time.Now()
	testLocationText = "testlocation"
)

func TestGetPosts(t *testing.T) {
	tests := []struct {
		name        string
		mockReturn  []sqlc.GetPostsRow
		mockErr     error
		expectPosts []*models.Posts
		expectedErr error
	}{
		{
			name: "Success - Returns posts",
			mockReturn: []sqlc.GetPostsRow{
				{PostID: testUUID, PostedDate: testTime, Latitude: &lat, Longtitude: &lng, AddressText: "test", Location: []byte(testLocationText)},
				{PostID: testUUID, PostedDate: testTime.Add(10 * time.Minute), Latitude: &lat, Longtitude: &lng, AddressText: "test2", Location: []byte(testLocationText)}},
			mockErr: nil,
			expectPosts: []*models.Posts{
				{PostID: testUUID, PostedDate: testTime, Latitude: &lat, Longtitude: &lng, AddressText: "test", Location: []byte(testLocationText)},
				{PostID: testUUID, PostedDate: testTime.Add(10 * time.Minute), Latitude: &lat, Longtitude: &lng, AddressText: "test2", Location: []byte(testLocationText)}},
			expectedErr: nil,
		},
		{
			name:        "Error - Failed to query",
			mockReturn:  nil,
			mockErr:     assert.AnError,
			expectPosts: nil,
			expectedErr: assert.AnError,
		},
		{
			name:        "Success - No posts found (empty slice)",
			mockReturn:  []sqlc.GetPostsRow{},
			mockErr:     nil,
			expectPosts: nil,
			expectedErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockRepo := repositories.NewMockPostRepository(t)
			mockRepo.On("GetPosts", mock.Anything).Return(tc.mockReturn, tc.mockErr)

			postService := services.NewPostService(mockRepo)
			result, err := postService.GetPosts(context.Background())

			if (err != nil && tc.expectedErr == nil) || (err == nil && tc.expectedErr != nil) || (err != nil && err.Error() != tc.expectedErr.Error()) {
				t.Errorf("예상 에러: '%v', 실제 에러: '%v'", tc.expectedErr, err)
			}
			if len(result) != len(tc.expectPosts) {
				t.Errorf("반환된 게시물 수가 다름. 예상: %d, 실제: %d", len(tc.expectPosts), len(result))
			}
		})
	}
}

func TestGetPostInfo(t *testing.T) {
	testDate := time.Date(2025, 9, 24, 0, 0, 0, 0, time.UTC)
	pgDate := pgtype.Date{
		Time:  testDate,
		Valid: true,
	}

	tests := []struct {
		name        string
		mockReturn  sqlc.GetPostInfoRow
		mockErr     error
		expectPost  *models.Post
		expectedErr error
	}{
		{
			name:        "Success - Return post info",
			mockReturn:  sqlc.GetPostInfoRow{PostID: testUUID, Content: "test-content", IncidentDate: pgDate, PostedDate: testTime, Latitude: &lat, Longtitude: &lng, AddressText: "test", Location: []byte(testLocationText)},
			mockErr:     nil,
			expectPost:  &models.Post{PostID: testUUID, Content: "test-content", IncidentDate: testDate, PostedDate: testTime, Latitude: &lat, Longtitude: &lng, AddressText: "test", Location: []byte(testLocationText)},
			expectedErr: nil,
		},
		{
			name:        "Error - Failed to query",
			mockReturn:  sqlc.GetPostInfoRow{},
			mockErr:     assert.AnError,
			expectPost:  &models.Post{},
			expectedErr: assert.AnError,
		},
		{
			name:        "Success - No post found (empty row)",
			mockReturn:  sqlc.GetPostInfoRow{},
			mockErr:     nil,
			expectPost:  &models.Post{},
			expectedErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockRepo := repositories.NewMockPostRepository(t)
			mockRepo.On("GetPostInfo", mock.Anything, testUUID).Return(tc.mockReturn, tc.mockErr)

			ctx := context.Background()

			postService := services.NewPostService(mockRepo)
			result, err := postService.GetPostInfo(ctx, testUUID.String())

			if (err != nil && tc.expectedErr == nil) || (err == nil && tc.expectedErr != nil) || (err != nil && err.Error() != tc.expectedErr.Error()) {
				t.Errorf("예상 에러: '%v', 실제 에러: '%v'", tc.expectedErr, err)
			}

			if tc.name == "Success - Return post info" {
				if !reflect.DeepEqual(result, tc.expectPost) {
					t.Errorf("예상 결과: '%v', 실제 결과: '%v'", tc.expectPost, result)
				}
			}
		})
	}
}

func TestCreatePost(t *testing.T) {
	var invalidTime time.Time
	tests := []struct {
		name        string
		inputValue  *models.Post
		expectedErr error
		mockErr     error
	}{
		{
			name:        "Success - Create a new post",
			inputValue:  &models.Post{Content: "create post test", IncidentDate: testTime, Latitude: &lat, Longtitude: &lng, AddressText: "test address"},
			expectedErr: nil,
			mockErr:     nil,
		},
		{
			name:        "Error - Failed to create a new post(Validation Check - empty space(content))",
			inputValue:  &models.Post{Content: "", IncidentDate: testTime, Latitude: &lat, Longtitude: &lng, AddressText: "test address"},
			expectedErr: errdefs.ErrEmptySpace,
			mockErr:     errdefs.ErrEmptySpace,
		},
		{
			name:        "Error - Failed to create a new post(Validation Check - empty space(incident_date))",
			inputValue:  &models.Post{Content: "create post test", IncidentDate: invalidTime, Latitude: &lat, Longtitude: &lng, AddressText: "test address"},
			expectedErr: errdefs.ErrEmptySpace,
			mockErr:     errdefs.ErrEmptySpace,
		},
		{
			name:        "Error - Failed to create a new post(Validation Check - empty space(latitude))",
			inputValue:  &models.Post{Content: "create post test", IncidentDate: testTime, Latitude: nil, Longtitude: &lng, AddressText: "test address"},
			expectedErr: errdefs.ErrEmptySpace,
			mockErr:     errdefs.ErrEmptySpace,
		},
		{
			name:        "Error - Failed to create a new post(Validation Check - empty space(longtitude))",
			inputValue:  &models.Post{Content: "create post test", IncidentDate: testTime, Latitude: &lat, Longtitude: nil, AddressText: "test address"},
			expectedErr: errdefs.ErrEmptySpace,
			mockErr:     errdefs.ErrEmptySpace,
		},
		{
			name:        "Error - Failed to create a new post(Validation Check - future time data)",
			inputValue:  &models.Post{Content: "create post test", IncidentDate: testTime.Add(5 * time.Second), Latitude: &lat, Longtitude: &lng, AddressText: "test address"},
			expectedErr: errdefs.ErrInvalidValue,
			mockErr:     errdefs.ErrInvalidValue,
		},
		{
			name:        "Error - Failed to create a new post(DB connection fail)",
			inputValue:  &models.Post{Content: "create post test", IncidentDate: testTime, Latitude: &lat, Longtitude: &lng, AddressText: "test address"},
			expectedErr: assert.AnError,
			mockErr:     assert.AnError,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockRepo := repositories.NewMockPostRepository(t)
			mockRepo.On("CreatePost", mock.Anything, mock.AnythingOfType("sqlc.CreatePostParams")).Return(tc.mockErr)

			ctx := context.Background()

			postService := services.NewPostService(mockRepo)
			err := postService.CreatePost(ctx, tc.inputValue)

			if (err != nil && tc.expectedErr == nil) || (err == nil && tc.expectedErr != nil) || (err != nil && tc.expectedErr != nil && !errors.Is(err, tc.expectedErr)) {
				t.Errorf("예상 에러: '%v', 실제 에러: '%v'", tc.expectedErr, err)
			}
		})
	}
}
