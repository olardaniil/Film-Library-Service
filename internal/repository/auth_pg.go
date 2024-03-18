package repository

import (
	"film-library-service/internal/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPG struct {
	db *sqlx.DB
}

func NewAuthPG(db *sqlx.DB) *AuthPG {
	return &AuthPG{db: db}
}

func (r *AuthPG) GetUser(username string, passwordHash string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id FROM users WHERE username=$1 AND password_hash=$2")
	err := r.db.Get(&user, query, username, passwordHash)

	return user, err
}
