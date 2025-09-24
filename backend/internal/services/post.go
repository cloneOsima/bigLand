package services

// Service package for implementing business logic

import (
	"context"
	"fmt"
	"time"

	"github.com/cloneOsima/bigLand/backend/internal/models"
	"github.com/cloneOsima/bigLand/backend/internal/repositories"
	"github.com/cloneOsima/bigLand/backend/internal/utils"
	"github.com/google/uuid"
)

type PostService interface {
	GetPosts(ctx context.Context) ([]*models.Posts, error)
	GetPostInfo(ctx context.Context) (*models.Post, error)
}

type postServiceImpl struct {
	postRepo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) PostService {
	return &postServiceImpl{postRepo: repo}
}

func (p *postServiceImpl) GetPosts(ctx context.Context) ([]*models.Posts, error) {

	// database context should be shorter than formal response context.
	// make new context
	dbCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// get data by using sqlc struct
	sqlcPosts, err := p.postRepo.GetPosts(dbCtx)
	if err != nil {
		return nil, err
	}

	// mapping sqlc struct <> models package
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

func (p *postServiceImpl) GetPostInfo(ctx context.Context) (*models.Post, error) {

	var postIdKey utils.CtxKey = "postId"

	// 전달 받은 UUID 값 valid test
	originValue := ctx.Value(postIdKey)
	strValue, ok := originValue.(string)
	if !ok {
		return nil, fmt.Errorf("error - given id has an invalid format")
	}
	ctxValue, parseErr := uuid.Parse(strValue)
	if parseErr != nil {
		return nil, parseErr
	}

	dbCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	dbCtx = context.WithValue(dbCtx, postIdKey, ctxValue)
	defer cancel()

	sqlcPost, err := p.postRepo.GetPostInfo(dbCtx)
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
