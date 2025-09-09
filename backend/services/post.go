package services

// Service package for implementing business logic

import (
	"context"
	"time"

	"github.com/cloneOsima/bigLand/backend/models"
	"github.com/cloneOsima/bigLand/backend/repositories"
)

type PostService interface {
	GetPosts(ctx context.Context) ([]models.EntirePost, error)
}

type postServiceImpl struct {
	postRepo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) PostService {
	return &postServiceImpl{postRepo: repo}
}

func (p *postServiceImpl) GetPosts(ctx context.Context) ([]models.EntirePost, error) {

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
