package repositories

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"

	"github.com/cloneOsima/bigLand/backend/models"
)

func ReadRegistInfo() {
	// load .env file
	err := godotenv.Load("../configs/postgresql.env")
	if err != nil {
		log.Fatal("Error loading postgresql.env file")
	}

	// create connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer conn.Close(context.Background())

	// a query for every row
	query := `
		SELECT
			id,
			accident_date,
			accident_cause,
			registartion_date,
			registed_addres
		FROM
			regist_info
	`

	// use .Query function for read every rows (single row needs .QueryRow())
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	var posts []models.Post

	// Scan
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Id,
			&post.AccidentDate,
			&post.AccidentCause,
			&post.RegistrationDate,
			&post.RegistratedAddress,
		)
		if err != nil {
			log.Fatalf("Scan failed: %v", err)
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Row iteration error: %v", err)
	}

	// test printing
	for _, post := range posts {
		fmt.Printf("Fetched data: %+v\n", post)
	}
}
