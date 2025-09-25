package services

// Service package for implementing business logic

import (
	"context"
	"time"

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
	// make new context
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

	// 전달 받은 string 값 uuid valid test 후 UUID로 변환
	postUUID, parseErr := uuid.Parse(postID)
	if parseErr != nil {
		return nil, parseErr
	}

	dbCtx, cancel := context.WithTimeout(ctx, defaultDBTimeout)
	defer cancel()

	sqlcPost, err := p.postRepo.GetPostInfo(dbCtx, postUUID)
	if err != nil {
		return nil, err
	}

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

	// generate db context
	dbCtx, cancel := context.WithTimeout(ctx, defaultDBTimeout)
	defer cancel()

	// data mapping
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
