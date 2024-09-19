package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Gierdiaz/diagier-clinics/internal/domain/user"
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
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
        return
    }
    
    err := h.service.Register(c.Request.Context(), req.Username, req.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register"})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *UserHandler) Login(c *gin.Context) {
    var req user.AuthRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
        return
    }
    
    token, err := h.service.Authenticate(c.Request.Context(), req.Username, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
        return
    }
    
    c.JSON(http.StatusOK, user.AuthResponse{Token: token})
}
