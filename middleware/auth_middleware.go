package middleware

import (
	"net/http"
	"strings"

	"mini-project-pbi/models"
	"mini-project-pbi/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token from header 'token' or 'Authorization'
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("token")
		if tokenHeader == "" {
			tokenHeader = c.GetHeader("Authorization")
		}

		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized", "Token tidak ditemukan"))
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")
		tokenString = strings.TrimSpace(tokenString)

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized", "Token tidak valid"))
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}
