package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/SaH1l191/ticket-system/internal/auth"
	"github.com/SaH1l191/ticket-system/internal/model"
	"github.com/SaH1l191/ticket-system/internal/repo"
	"github.com/google/uuid"
)

var (
	ErrEmailExists      = errors.New("email already registered")
	ErrInvalidCredentials = errors.New("invalid email or password")
)

type AuthService struct {
	users repo.UserRepository
}

func NewAuthService(u repo.UserRepository) *AuthService {
	return &AuthService{users: u}
}

func (s *AuthService) Register(ctx context.Context, email, password string) error {
	email = strings.TrimSpace(email)
	if email == "" || password == "" {
		return errors.New("email and password are required")
	}

	existing, err := s.users.GetByEmail(ctx, email)
	if err != nil {
		return err
	}
	if existing != nil {
		return ErrEmailExists
	}

	hash, err := auth.HashPassword(password)
	if err != nil {
		return err
	}

	user := model.User{
		ID:           uuid.New().String(),
		Email:        email,
		PasswordHash: hash,
		CreatedAt:    time.Now(),
	}

	return s.users.Create(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	email = strings.TrimSpace(email)
	if email == "" || password == "" {
		return "", ErrInvalidCredentials
	}

	user, err := s.users.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if user == nil || !auth.CheckPassword(user.PasswordHash, password) {
		return "", ErrInvalidCredentials
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
