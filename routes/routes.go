package routes

import (
	"mini-project-pbi/handlers"
	"mini-project-pbi/middleware"
	"mini-project-pbi/repositories"
	"mini-project-pbi/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes registers all application routes and dependency injections
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Static folder for uploaded files
	router.Static("/uploads", "./uploads")

	// Repositories
	userRepo := repositories.NewUserRepository(db)
	tokoRepo := repositories.NewTokoRepository(db)

	// Services
	authService := services.NewAuthService(userRepo, tokoRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)

	// Auth Routes (Public)
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Protected User Ping (for verifying Auth Middleware)
	userProtected := router.Group("/user", middleware.AuthMiddleware())
	{
		userProtected.GET("/ping", func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			c.JSON(200, gin.H{
				"status":  true,
				"message": "Auth middleware working",
				"user_id": userID,
			})
		})
	}
}
