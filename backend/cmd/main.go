package main

// Cmd package for using server's entry point.

import (
	"fmt"
	"log"

	"github.com/cloneOsima/bigLand/backend/server"
)

func main() {
	port := 10001
	r := server.SetupRouter()

	addr := fmt.Sprintf(":%d", port) // ":10001" 형태로 문자열 생성
	log.Printf("Server starting on http://localhost:%d", port)
	r.Run(addr)
}
