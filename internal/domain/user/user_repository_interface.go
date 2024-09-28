package user

import (
	"context"
)

type UserRepository interface {
	Email(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) error
}
