package main

import (
	"fmt"

	"github.com/cloneOsima/bigLand/backend/handlers"
	"github.com/cloneOsima/bigLand/backend/repositories"
	"github.com/cloneOsima/bigLand/backend/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

// 서버 셋업 시 각 레이어 별 interface 조립
func setUp(pool *pgxpool.Pool) (handlers.Handler, error) {

	// repo 생성
	postRepo := repositories.NewPostRepository(pool)

	// service 생성
	postSvc := services.NewPostService(postRepo)

	// handler 생성
	handler := handlers.NewHandler(postSvc)
	if handler == nil {
		return nil, fmt.Errorf("failed to create handler")
	}

	return handler, nil
}
