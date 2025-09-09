package main

// Cmd package for using server's entry point.

import (
	"fmt"
	"log"

	"github.com/cloneOsima/bigLand/backend/server"
)

func main() {
	port := 10001
	// Initiate services code

	// Create DB connection pool
	dbPool, err := initConnectionPool()
	if err != nil {
		log.Fatalf("Failed to initialize connection pool: %v", err)
	}
	defer deleteConnectionPool()
	log.Printf("Passed: Connection pool is created.")

	// Setup
	handler, err := setUp(dbPool)
	if err != nil {
		log.Fatalf("Failed to setup interfaces: %v", err)
	}

	r := server.SetupRouter(handler)

	addr := fmt.Sprintf(":%d", port) // ":10001" 형태로 문자열 생성
	log.Printf("Server starting on http://localhost:%d", port)
	r.Run(addr)
}
