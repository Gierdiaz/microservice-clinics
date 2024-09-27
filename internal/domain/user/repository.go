package user

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Email(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) Repository {
	return &userRepository{db: db}
}

func (r *userRepository) Email(ctx context.Context, email string) (*User, error) {
	var user User
	err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, user *User) error {
	_, err := r.db.NamedExecContext(ctx, `
		INSERT INTO users (id, name, email, password, created_at, updated_at) 
		VALUES (:id, :name, :email, :password, NOW(), NOW())`, user)
	return err
}
