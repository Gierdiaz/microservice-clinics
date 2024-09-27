package middleware

import (
	"github.com/Gierdiaz/diagier-clinics/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

var cfg *config.Config

func InitJWT(config *config.Config) {
	cfg = config
}

func GenerateToken(userID uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(cfg.JWT.ExpHours)).Unix()

	tokenString, err := token.SignedString([]byte(cfg.JWT.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, nil, err
	}
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, err
	}
	return token, claims, nil
}
