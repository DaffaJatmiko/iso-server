package middleware

import (
	"github.com/DaffaJatmiko/project-iso/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is empty"})
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(authHeader)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}

		claims, err := util.ParseJWT(tokenString, jwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()

	}
}
