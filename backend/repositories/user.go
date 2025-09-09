package repositories

import (
	"github.com/cloneOsima/bigLand/backend/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Login()
	Logout()
	NewAccount()
	DeleteAccount()
}

type User struct {
	dbPool   *pgxpool.Pool
	UserData models.User
}

func (u *User) Login() error {

	return nil
}

func (u *User) Logout() error {

	return nil
}

func (u *User) NewAccount() error {

	return nil
}

func (u *User) DeleteAccount() error {

	return nil
}
