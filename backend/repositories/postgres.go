package repositories

// Repositories package for accessing the database and performing CRUD operations

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/cloneOsima/bigLand/backend/models"
)

var dbPool *pgxpool.Pool

func InitDB() error {
	err := godotenv.Load("configs/postgresql.env")
	if err != nil {
		log.Printf("Errors: Failed to load postgresql.env file")
		return err
	}

	// Connection pool 설정
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Printf("Errors: Failed to parse database URL: %v", err)
		return err
	}

	// Connection pool 설정값 조정
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = time.Minute * 30
	config.HealthCheckPeriod = time.Minute * 5

	// Connection pool 생성
	dbPool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Printf("Errors: Failed to create connection pool: %v", err)
		return err
	}

	// 연결 테스트
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = dbPool.Ping(ctx)
	if err != nil {
		log.Printf("Errors: Failed to ping database: %v", err)
		return err
	}

	log.Println("Successfully connected to Supabase database")
	return nil
}

func CloseDB() {
	if dbPool != nil {
		dbPool.Close()
	}
}

func GetEntirePost(dbCtx context.Context) ([]models.EntirePost, error) {
	if dbPool == nil {
		return nil, fmt.Errorf("database pool is not initialized")
	}

	query := `
		SELECT
			id,
			created_at,
			registed_addres
		FROM
			post
	`

	// use .Query function for read every rows (single row needs .QueryRow())
	rows, err := dbPool.Query(dbCtx, query)
	if err != nil {
		log.Printf("Errors: Query failed: %v", err)
		return nil, err
	}
	defer rows.Close()

	var posts []models.EntirePost

	// Scan
	for rows.Next() {
		var post models.EntirePost
		err := rows.Scan(
			&post.Id,
			&post.CreatedAt,
			&post.RegistratedAddress,
		)
		if err != nil {
			log.Printf("Errors: Scan failed: %v", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Errors: Row iteration error: %v", err)
		return nil, err
	}

	// test printing
	for _, post := range posts {
		fmt.Printf("Errors: Fetched data: %+v\n", post)
	}

	return posts, nil
}
