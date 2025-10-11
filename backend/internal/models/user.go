package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId      uuid.UUID `db:"user_id" json:"user_id"`
	Username    string    `db:"username" json:"username"`
	Email       string    `db:"email" json:"email"`
	Password    string    `db:"password_hash" json:"password_hash"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	LastLoginAt time.Time `db:"last_login_at" json:"last_login_at"`
	IsActive    bool      `db:"is_Active" json:"is_Active"`
}
