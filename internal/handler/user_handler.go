package handler

import (
	"fmt"
	"net/http"

	"github.com/Gierdiaz/diagier-clinics/internal/domain/user"
	"github.com/Gierdiaz/diagier-clinics/pkg/validator"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req user.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	if validationErrors := validator.Validate(req); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "errors": validationErrors})
		return
	}

	err := h.service.Register(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req user.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Erro ao fazer bind do JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	fmt.Println("Tentando autenticar com o email:", req.Email)

	token, err := h.service.Authenticate(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		fmt.Println("Erro de autenticação:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	fmt.Println("Login bem-sucedido, retornando token.")
	c.JSON(http.StatusOK, user.AuthResponse{Token: token})
}
