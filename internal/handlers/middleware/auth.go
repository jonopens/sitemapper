package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates authentication tokens
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement authentication logic
		// 1. Extract token from header
		// 2. Validate token
		// 3. Set user context
		// 4. Continue or abort
		
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Missing authorization token",
			})
			c.Abort()
			return
		}
		
		// Validate token here
		
		c.Next()
	}
}

