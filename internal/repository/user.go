package repository

import "github.com/Coldwws/todo/internal/models"

type UserRepository interface {
	GetByUsername(username string) (*models.User, error)
	Create(username, passwordHash string) error
}
