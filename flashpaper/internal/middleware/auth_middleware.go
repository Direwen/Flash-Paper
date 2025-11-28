package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/direwen/flashpaper/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the auth header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.SendError(c, http.StatusUnauthorized, errors.New("authorization header is required"))
			c.Abort()
			return
		}

		// Parse the format "Bearer/JWT <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.SendError(c, http.StatusUnauthorized, errors.New("invalid authorization header format"))
			c.Abort()
			return
		}
		// Validate Token
		tokenString := parts[1]
		userID, err := utils.ValidateToken(tokenString)
		if err != nil {
			utils.SendError(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		// Store User ID in Context
		c.Set("userID", userID)

		// Next Handler
		c.Next()
	}
}
