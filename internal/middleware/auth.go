package middleware

import (
	"go-crud-msql/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or malformed token"})
			return
		}

		// Extract token string
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		tokenString = strings.TrimSpace(tokenString)

		userData, err := utils.ValidateJWT(tokenString)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errMessage": err.Error(), "error": "Invalid or expired token"})
			return
		}

		// Store userData in context
		c.Set("userData", userData)
		c.Next()
	}
}
