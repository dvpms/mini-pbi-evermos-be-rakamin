package routes

import (
	"mini-project-pbi/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes registers public authentication endpoints
func RegisterAuthRoutes(router *gin.Engine, authHandler *handlers.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}
}
