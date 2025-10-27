package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "mysecrettoken" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "未授权"})
			c.Abort()
			return
		}
		c.Next()
	}
}
