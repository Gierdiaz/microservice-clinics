package user

import "github.com/google/uuid"

type UserDTO struct {
	ID       uuid.UUID `json:"id"`
	Name     string `json:"name" validate:"required"`
}

type AuthRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=3"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
