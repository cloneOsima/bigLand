package repositories

// Repositories package for accessing the database and performing CRUD operations

import (
	"context"

	"github.com/cloneOsima/bigLand/backend/internal/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository interface {
	GetPosts(ctx context.Context) ([]sqlc.GetPostsRow, error)
	GetPostInfo(ctx context.Context, postID uuid.UUID) (sqlc.GetPostInfoRow, error)
	CreatePost(dbCtx context.Context, info sqlc.CreatePostParams) error
}

type postRepoImpl struct {
	q *sqlc.Queries
}

func NewPostRepository(pool *pgxpool.Pool) PostRepository {
	return &postRepoImpl{
		q: sqlc.New(pool),
	}
}

func (p *postRepoImpl) GetPosts(dbCtx context.Context) ([]sqlc.GetPostsRow, error) {
	return p.q.GetPosts(dbCtx)
}

func (p *postRepoImpl) GetPostInfo(dbCtx context.Context, postID uuid.UUID) (sqlc.GetPostInfoRow, error) {
	return p.q.GetPostInfo(dbCtx, postID)
}

func (p *postRepoImpl) CreatePost(dbCtx context.Context, info sqlc.CreatePostParams) error {
	return p.q.CreatePost(dbCtx, info)
}

// type PostRepository interface {
// 	GetPosts(ctx context.Context) ([]*models.Posts, error)
// 	GetPostInfo(ctx context.Context) (*models.Post, error)
// }

// type postRepoImpl struct {
// 	dbPool *pgxpool.Pool
// }

// func NewPostRepository(pool *pgxpool.Pool) PostRepository {
// 	return &postRepoImpl{
// 		dbPool: pool,
// 	}
// }

// func (p *postRepoImpl) GetPosts(dbCtx context.Context) ([]*models.Posts, error) {
// 	if p.dbPool == nil {
// 		return nil, fmt.Errorf("connection pool is not initialized")
// 	}

// 	query := `
// 		SELECT post_id, posted_date, address_text, latitude, longtitude, location
// 		FROM posts
// 		WHERE is_active = true
// 		ORDER BY posted_date DESC
// 		LIMIT 50;
// 	`

// 	// .Query function 다중 검색을 위한 사용 (단일행 검색은 .QueryRow()사용)
// 	rows, err := p.dbPool.Query(dbCtx, query)
// 	if err != nil {
// 		log.Printf("Errors: Query failed: %v", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var result = make([]*models.Posts, 0, 50)

// 	// 결과 row 스캔 및 dataset에 매핑
// 	for rows.Next() {
// 		post := &models.Posts{}
// 		err := rows.Scan(
// 			&post.PostId,
// 			&post.PostedDate,
// 			&post.AddressText,
// 			&post.Latitude,
// 			&post.Longtitude,
// 			&post.Location,
// 		)
// 		if err != nil {
// 			log.Printf("Errors: Scan failed: %v", err)
// 			return nil, err
// 		}
// 		result = append(result, post)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("Errors: Row iteration error: %v", err)
// 		return nil, err
// 	}

// 	// test printing
// 	for _, post := range result {
// 		fmt.Printf("Fetched data: %+v\n", post)
// 	}

// 	return result, nil
// }

// func (p *postRepoImpl) GetPostInfo(dbCtx context.Context) (*models.Post, error) {
// 	if p.dbPool == nil {
// 		return nil, fmt.Errorf("connection pool is not initialized")
// 	}
// 	var postIdKey utils.CtxKey = "postId"
// 	postId := dbCtx.Value(postIdKey)

// 	query := `
// 		SELECT post_id, content, incident_date, posted_date, address_text, latitude, longtitude, location
// 		FROM posts
// 		WHERE is_active = true
// 		AND post_id = $1;
// 	`

// 	row := p.dbPool.QueryRow(dbCtx, query, postId)
// 	result := &models.Post{}

// 	err := row.Scan(
// 		&result.PostId,
// 		&result.Content,
// 		&result.IncidentDate,
// 		&result.PostedDate,
// 		&result.AddressText,
// 		&result.Latitude,
// 		&result.Longtitude,
// 		&result.Location,
// 	)
// 	if err != nil {
// 		if err == pgx.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return result, nil
// }
