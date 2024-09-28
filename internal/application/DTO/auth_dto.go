package DTO

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
