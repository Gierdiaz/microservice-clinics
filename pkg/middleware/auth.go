package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token is required"})
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

		// Converte o ID para uuid.UUID
		parsedID, err := uuid.Parse(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user ID format"})
			c.Abort()
			return
		}

		// Adiciona o ID do usuário no contexto para uso posterior
		c.Set("userID", parsedID)
		c.Next()
	}
}
