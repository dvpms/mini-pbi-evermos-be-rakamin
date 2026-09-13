package routes

import (
	"mini-project-pbi/handlers"
	"mini-project-pbi/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registers user profile and address management endpoints
func RegisterUserRoutes(router *gin.Engine, userHandler *handlers.UserHandler, alamatHandler *handlers.AlamatHandler) {
	user := router.Group("/user", middleware.AuthMiddleware())
	{
		// Profil User
		user.GET("", userHandler.GetProfile)
		user.PUT("", userHandler.UpdateProfile)

		// Alamat Kirim
		user.GET("/alamat", alamatHandler.GetMyAlamat)
		user.GET("/alamat/:id", alamatHandler.GetAlamatByID)
		user.POST("/alamat", alamatHandler.CreateAlamat)
		user.PUT("/alamat/:id", alamatHandler.UpdateAlamat)
		user.DELETE("/alamat/:id", alamatHandler.DeleteAlamat)

		// Ping test
		user.GET("/ping", func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			c.JSON(200, gin.H{
				"status":  true,
				"message": "Auth middleware working",
				"user_id": userID,
			})
		})
	}
}
