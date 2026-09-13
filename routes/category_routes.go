package routes

import (
	"mini-project-pbi/handlers"
	"mini-project-pbi/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterCategoryRoutes registers category retrieval and admin mutation endpoints
func RegisterCategoryRoutes(router *gin.Engine, categoryHandler *handlers.CategoryHandler) {
	// Public category routes
	category := router.Group("/category")
	{
		category.GET("", categoryHandler.GetAllCategories)
		category.GET("/:id", categoryHandler.GetCategoryByID)
	}

	// Admin only category mutation routes
	categoryAdmin := router.Group("/category", middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		categoryAdmin.POST("", categoryHandler.CreateCategory)
		categoryAdmin.PUT("/:id", categoryHandler.UpdateCategory)
		categoryAdmin.DELETE("/:id", categoryHandler.DeleteCategory)
	}
}
