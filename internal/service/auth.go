package service

import (
	"errors"

	"github.com/Coldwws/todo/internal/auth"
	"github.com/Coldwws/todo/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	users repository.UserRepository
}

func NewAuthService(users repository.UserRepository) AuthService {
	return &authService{users: users}
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.users.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	); err != nil {
		return "", errors.New("invalid credentials")
	}

	return auth.GenerateToken(user.ID)
}

func (s *authService) Register(username, password string) error {
	_, err := s.users.GetByUsername(username)
	if err == nil{
		return errors.New("username already exists")
	}
	hash,err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.users.Create(username, string(hash))
}
