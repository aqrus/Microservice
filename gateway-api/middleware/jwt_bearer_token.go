package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtBearerToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Verify JWT token here

		c.Next()
	}
}
