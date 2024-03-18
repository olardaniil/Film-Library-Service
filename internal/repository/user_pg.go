package repository

import (
	"film-library-service/internal/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPG struct {
	db *sqlx.DB
}

func NewUserPG(db *sqlx.DB) *UserPG {
	return &UserPG{db: db}
}

func (r *UserPG) IsUserAdmin(userID int) error {
	var user entity.User
	query := fmt.Sprintf("SELECT id FROM users WHERE id=$1 AND is_admin=$2")
	err := r.db.Get(&user, query, userID, true)
	if err != nil {
		return err
	}
	return nil
}
