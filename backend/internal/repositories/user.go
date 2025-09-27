package repositories

import (
	"context"

	"github.com/cloneOsima/bigLand/backend/internal/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Login() error
	Logout() error
	NewAccount(dbCtx context.Context, newAccount sqlc.InsertNewAccountParams) (sqlc.SelectUserRow, error)
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

func (u *userRepoImpl) NewAccount(dbCtx context.Context, newAccount sqlc.InsertNewAccountParams) (sqlc.SelectUserRow, error) {
	// user account add
	err := u.q.InsertNewAccount(dbCtx, newAccount)
	if err != nil {
		return sqlc.SelectUserRow{}, err
	}

	// get added account from db
	return u.q.SelectUser(dbCtx, newAccount.Username)
}

func (u *userRepoImpl) DeleteAccount() error {

	return nil
}
