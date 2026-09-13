package middleware

import (
	"net/http"

	"mini-project-pbi/models"

	"github.com/gin-gonic/gin"
)

// AdminOnly verifies that authenticated user has admin role
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized", "Hanya admin yang memiliki akses"))
			c.Abort()
			return
		}
		c.Next()
	}
}
