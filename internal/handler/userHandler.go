package handler

import (
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
    
    // Valida os dados de registro usando o validador específico
    if validationErrors := validator.ValidateRegister(req); validationErrors != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "errors": validationErrors})
        return
    }

    // Se os dados forem válidos, prossegue com o registro
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
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
        return
    }
    
    token, err := h.service.Authenticate(c.Request.Context(), req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
        return
    }
    
    c.JSON(http.StatusOK, user.AuthResponse{Token: token})
}
