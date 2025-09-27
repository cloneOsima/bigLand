package main

// Cmd package for using server's entry point.

import (
	"fmt"
	"log"

	"github.com/cloneOsima/bigLand/backend/internal/handlers"
	"github.com/cloneOsima/bigLand/backend/internal/repositories"
	"github.com/cloneOsima/bigLand/backend/internal/server"
	"github.com/cloneOsima/bigLand/backend/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	portNo = 10001
)

func main() {
	// Initiate services code

	// Create DB connection pool
	dbPool, err := repositories.InitPool()
	if err != nil {
		log.Fatalf("Failed to initialize connection pool: %v", err)
	}
	defer repositories.DropPool()
	log.Printf("Passed: Connection pool is created.")

	// Setup
	handler, err := setUp(dbPool)
	if err != nil {
		log.Fatalf("Failed to setup interfaces: %v", err)
	}

	r := server.SetupRouter(handler)

	addr := fmt.Sprintf(":%d", portNo) // ":10001" 형태로 문자열 생성
	log.Printf("Server starting on http://localhost:%d", portNo)
	r.Run(addr)
}

func setUp(pool *pgxpool.Pool) (*handlers.Handlers, error) {

	// repo 초기화
	repo := repositories.InitRepo(pool)

	// service 초기화
	svc := services.InitSvc(repo)

	// handler 초기화
	handler := handlers.InitHandler(svc)
	if handler == nil {
		return nil, fmt.Errorf("failed to create handler")
	}

	return handler, nil
}
