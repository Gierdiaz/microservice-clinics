package user

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindByUsername(ctx context.Context, username string) (*User, error)
	Create(ctx context.Context, user *User) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) Repository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, user *User) error {
	_, err := r.db.NamedExecContext(ctx, "INSERT INTO users (id, name, email, password, role) VALUES (:id, :name, :email, :password, :role)", user)
	return err
}
