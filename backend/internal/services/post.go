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

	result, err := p.postRepo.GetPosts(dbCtx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *postServiceImpl) GetPostInfo(ctx context.Context) (*models.Post, error) {

	var postIdKey utils.CtxKey = "postId"

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

	result, err := p.postRepo.GetPostInfo(dbCtx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
