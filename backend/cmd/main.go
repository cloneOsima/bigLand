package main

// Cmd package for using server's entry point.

import (
	"fmt"
	"log"

	"github.com/cloneOsima/bigLand/backend/repositories"
	"github.com/cloneOsima/bigLand/backend/server"
)

func main() {
	port := 10001
	// Initiate services code

	// Create DB connection pool
	err := repositories.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer repositories.CloseDB()

	r := server.SetupRouter()

	addr := fmt.Sprintf(":%d", port) // ":10001" 형태로 문자열 생성
	log.Printf("Server starting on http://localhost:%d", port)
	r.Run(addr)
}
