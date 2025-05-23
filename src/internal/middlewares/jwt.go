package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckAuthorizationHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			return
		}

		// Continue to the next handler
		c.Next()
	}
}
