package user

import (
	"context"
	"github.com/Gierdiaz/diagier-clinics/pkg/middleware"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

type Service interface {
	Authenticate(ctx context.Context, username, password string) (string, error)
	Register(ctx context.Context, username, password string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Authenticate(ctx context.Context, username, password string) (string, error) {
	user, err := s.repo.FindByUsername(ctx, username)
	if err != nil || user == nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) Register(ctx context.Context, username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &User{
		ID:       uuid.New(),
		Name:     username,   
		Password: string(hashedPassword),
	}

	return s.repo.Create(ctx, user)
}
