package services

// Service package for implementing business logic

import (
	"context"
	"time"

	errdefs "github.com/cloneOsima/bigLand/backend/internal/errors"
	"github.com/cloneOsima/bigLand/backend/internal/models"
	"github.com/cloneOsima/bigLand/backend/internal/repositories"
	"github.com/cloneOsima/bigLand/backend/internal/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostService interface {
	GetPosts(ctx context.Context) ([]*models.Posts, error)
	GetPostInfo(ctx context.Context, postID string) (*models.Post, error)
	CreatePost(ctx context.Context, inputValue *models.Post) error
}

type postServiceImpl struct {
	postRepo repositories.PostRepository
}

const defaultDBTimeout = 5 * time.Second

func NewPostService(repo repositories.PostRepository) PostService {
	return &postServiceImpl{postRepo: repo}
}

func (p *postServiceImpl) GetPosts(ctx context.Context) ([]*models.Posts, error) {

	// database context should be shorter than formal response context.
	// make new db context
	dbCtx, cancel := context.WithTimeout(ctx, defaultDBTimeout)
	defer cancel()

	// get data by using sqlc struct
	sqlcPosts, err := p.postRepo.GetPosts(dbCtx)
	if err != nil {
		return nil, err
	}

	// mapping sqlc struct <> models package struct
	var result []*models.Posts
	for _, sp := range sqlcPosts {
		result = append(result, &models.Posts{
			PostID:      sp.PostID,
			PostedDate:  sp.PostedDate,
			AddressText: sp.AddressText,
			Latitude:    sp.Latitude,
			Longtitude:  sp.Longtitude,
			Location:    sp.Location,
		})
	}

	return result, nil
}

func (p *postServiceImpl) GetPostInfo(ctx context.Context, postID string) (*models.Post, error) {

	// validation check
	postUUID, parseErr := uuid.Parse(postID)
	if parseErr != nil {
		return nil, parseErr
	}

	// make new db context
	dbCtx, cancel := context.WithTimeout(ctx, defaultDBTimeout)
	defer cancel()

	// get data by using sqlc struct
	sqlcPost, err := p.postRepo.GetPostInfo(dbCtx, postUUID)
	if err != nil {
		return nil, err
	}

	// mapping sqlc struct <> models package struct
	var result = new(models.Post)
	result = &models.Post{
		PostID:       sqlcPost.PostID,
		Content:      sqlcPost.Content,
		IncidentDate: sqlcPost.IncidentDate.Time,
		PostedDate:   sqlcPost.PostedDate,
		Latitude:     sqlcPost.Latitude,
		Longtitude:   sqlcPost.Longtitude,
		AddressText:  sqlcPost.AddressText,
		Location:     sqlcPost.Location,
	}

	return result, nil
}

func (p *postServiceImpl) CreatePost(ctx context.Context, info *models.Post) error {

	// validation check
	valErr := createValueCheck(info)
	if valErr != nil {
		return valErr
	}

	// generate db context
	dbCtx, cancel := context.WithTimeout(ctx, defaultDBTimeout)
	defer cancel()

	// data mapping models package struct <> sqlc struct
	sqlcData := sqlc.CreatePostParams{
		Content:      info.Content,
		IncidentDate: pgtype.Date{Time: info.IncidentDate, Valid: true},
		Latitude:     info.Latitude,
		Longtitude:   info.Longtitude,
		AddressText:  info.AddressText,
	}

	// repo connection
	err := p.postRepo.CreatePost(dbCtx, sqlcData)
	if err != nil {
		return err
	}
	return nil
}

// CreatePost function validation check 가 너무 길어저셔 뺴놓은 함수
func createValueCheck(info *models.Post) *errdefs.AppError {
	var nTime time.Time
	switch {
	case info.Content == "":
		e := *errdefs.ErrEmptySpace
		e.ErrorInfo = []string{"Content"}
		return &e
	case info.IncidentDate.Equal(nTime):
		e := *errdefs.ErrEmptySpace
		e.ErrorInfo = []string{"IncidentDate"}
		return &e
	case info.Latitude == nil:
		e := *errdefs.ErrEmptySpace
		e.ErrorInfo = []string{"Latitude"}
		return &e
	case info.Longtitude == nil:
		e := *errdefs.ErrEmptySpace
		e.ErrorInfo = []string{"Longtitude"}
		return &e
	case info.IncidentDate.After(time.Now()):
		e := *errdefs.ErrInvalidValue
		e.ErrorInfo = []string{"IncidentDate", "future date"}
		return &e
	}
	return nil
}
