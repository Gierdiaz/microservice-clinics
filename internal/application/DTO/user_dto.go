package DTO

import "github.com/google/uuid"

type UserDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name" validate:"required"`
}