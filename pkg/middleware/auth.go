package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Token is required"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader { 
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
            c.Abort()
            return
        }

        _, claims, err := ParseToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
            c.Abort()
            return
        }

        userID, ok := (*claims)["id"].(string)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
            c.Abort()
            return
        }

        parsedID, err := uuid.Parse(userID)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user ID format"})
            c.Abort()
            return
        }

        c.Set("userID", parsedID)
        c.Next()
    }
}

