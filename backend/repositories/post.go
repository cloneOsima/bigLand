package repositories

// Repositories package for accessing the database and performing CRUD operations

import (
	"context"
	"fmt"
	"log"

	"github.com/cloneOsima/bigLand/backend/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository interface {
	GetPosts(ctx context.Context) ([]models.EntirePost, error)
}

type postRepoImpl struct {
	dbPool   *pgxpool.Pool
	PostData models.Post
}

func NewPostRepository(pool *pgxpool.Pool) PostRepository {
	return &postRepoImpl{
		dbPool: pool,
	}
}

func (p *postRepoImpl) GetPosts(dbCtx context.Context) ([]models.EntirePost, error) {
	if p.dbPool == nil {
		return nil, fmt.Errorf("connection pool is not initialized")
	}

	query := `
		SELECT
			post_id,
			posted_date,
			address_text
		FROM
			posts
	`

	// .Query function 다중 검색을 위한 사용 (단일행 검색은 .QueryRow()사용)
	rows, err := p.dbPool.Query(dbCtx, query)
	if err != nil {
		log.Printf("Errors: Query failed: %v", err)
		return nil, err
	}
	defer rows.Close()

	var posts []models.EntirePost

	// 결과 row 스캔 및 dataset에 매핑
	for rows.Next() {
		var post models.EntirePost
		err := rows.Scan(
			&post.PostId,
			&post.PostedDate,
			&post.AddressText,
		)
		if err != nil {
			log.Printf("Errors: Scan failed: %v", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Errors: Row iteration error: %v", err)
		return nil, err
	}

	// test printing
	for _, post := range posts {
		fmt.Printf("Fetched data: %+v\n", post)
	}

	return posts, nil
}
