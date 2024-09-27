package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gierdiaz/diagier-clinics/config"
	userApplication "github.com/Gierdiaz/diagier-clinics/internal/application/user"
	"github.com/Gierdiaz/diagier-clinics/internal/domain/user"
	"github.com/Gierdiaz/diagier-clinics/internal/handler"
	"github.com/Gierdiaz/diagier-clinics/pkg/middleware"
	"github.com/Gierdiaz/diagier-clinics/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockUserService struct {
	mockUsers map[string]*user.User
}

func newMockUserService() *mockUserService {
	return &mockUserService{mockUsers: make(map[string]*user.User)}
}

func (m *mockUserService) Authenticate(ctx context.Context, email, password string) (string, error) {
	usr, exists := m.mockUsers[email]
	if !exists || usr.Password != password {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := middleware.GenerateToken(usr.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (m *mockUserService) Register(ctx context.Context, email, password string) error {
	m.mockUsers[email] = &user.User{
		Email:    email,
		Password: password,
	}
	return nil
}

func TestRegister(t *testing.T) {
	validator.InitValidator()

	userService := newMockUserService()
	handler := handler.NewUserHandler(userService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/register", handler.Register)

	registerBody := userApplication.AuthRequest{
		Email:    "testuser@example.com",
		Password: "securepassword",
	}
	bodyBytes, _ := json.Marshal(registerBody)

	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	var response map[string]string
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "User registered successfully", response["message"])
}

func TestLogin(t *testing.T) {
	cfg := &config.Config{
		JWT: config.JWT{
			Secret:   "secrettoken",
			ExpHours: 1,
		},
	}

	middleware.InitJWT(cfg)

	userService := newMockUserService()
	handler := handler.NewUserHandler(userService)

	userService.Register(context.Background(), "testuser@example.com", "securepassword")

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", handler.Login)

	loginBody := userApplication.AuthRequest{
		Email:    "testuser@example.com",
		Password: "securepassword",
	}
	bodyBytes, _ := json.Marshal(loginBody)

	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response userApplication.AuthResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.NotEmpty(t, response.Token)
}
