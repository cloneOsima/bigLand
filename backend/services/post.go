package services

// Service package for implementing business logic

import (
	"github.com/cloneOsima/bigLand/backend/models"
	"github.com/cloneOsima/bigLand/backend/repositories"
)

func GetPostList() ([]models.EntirePost, error) {

	result, err := repositories.GetEntirePost()
	if err != nil {
		return nil, err
	}

	return result, nil
}
