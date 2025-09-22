package repositories

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var dbPool *pgxpool.Pool

// 서버 셋업 시 tracnsaction pool 생성
func InitPool() (*pgxpool.Pool, error) {
	err := godotenv.Load("internal/configs/postgresql.env")
	if err != nil {
		log.Fatalf("Errors: Failed to load postgresql.env file")
		return nil, err
	}

	// Connection pool 설정
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Errors: Failed to parse database URL: %v", err)
		return nil, err
	}

	// Connection pool 설정값 조정
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = time.Minute * 30
	config.HealthCheckPeriod = time.Minute * 5
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol // transaction pool 사용으로 인한 prepared statement 비활성화 설정

	// Connection pool 생성
	dbPool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Errors: Failed to create connection pool: %v", err)
		return nil, err
	}

	// 연결 테스트
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = dbPool.Ping(ctx)
	if err != nil {
		log.Fatalf("Errors: Failed to ping database: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to Supabase database")
	return dbPool, nil
}

// 서버 종료 시 connection pool 제거
func DropPool() {
	if dbPool != nil {
		dbPool.Close()
	}
}
