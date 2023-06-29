package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func FormatJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}
