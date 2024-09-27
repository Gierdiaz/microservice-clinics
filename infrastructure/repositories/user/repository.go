package user

import (
	"context"

	"github.com/Gierdiaz/diagier-clinics/internal/domain/user"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) user.UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) Email(ctx context.Context, email string) (*user.User, error) {
	var u user.User
	err := repo.db.GetContext(ctx, &u, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (repo *userRepository) Create(ctx context.Context, u *user.User) error {
	_, err := repo.db.NamedExecContext(ctx, `
		INSERT INTO users (id, name, email, password, created_at, updated_at) 
		VALUES (:id, :name, :email, :password, NOW(), NOW())`, u)
	return err
}
