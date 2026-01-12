package repository

import (
	"github.com/Coldwws/todo/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password_hash FROM users WHERE username=$1`
	err := r.db.Get(&user, query, username)
	return &user, err
}

func (r *UserPostgres) Create(username, passwordHash string) error {
	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2)`
	_, err := r.db.Exec(query, username, passwordHash)
	return err
}