package repositories

import (
	"context"

	"github.com/cloneOsima/bigLand/backend/internal/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Login() error
	Logout() error
	InsertNewAccount(dbCtx context.Context, newAccount sqlc.InsertNewAccountParams) error
	DeleteAccount() error
}

type userRepoImpl struct {
	q *sqlc.Queries
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return &userRepoImpl{
		q: sqlc.New(pool),
	}
}

func (u *userRepoImpl) Login() error {

	return nil
}

func (u *userRepoImpl) Logout() error {

	return nil
}

func (u *userRepoImpl) InsertNewAccount(dbCtx context.Context, newAccount sqlc.InsertNewAccountParams) error {
	// user account add
	err := u.q.InsertNewAccount(dbCtx, newAccount)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepoImpl) DeleteAccount() error {

	return nil
}
