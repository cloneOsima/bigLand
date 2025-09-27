// server 초기화 시 repositories 를 한번에 묶어서 초기화

package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type Repositories struct {
	User UserRepository
	Post PostRepository
}

func InitRepo(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		User: NewUserRepository(pool),
		Post: NewPostRepository(pool),
	}
}
