package routes

import (
	"mini-project-pbi/handlers"
	"mini-project-pbi/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterTokoRoutes registers store management endpoints
func RegisterTokoRoutes(router *gin.Engine, tokoHandler *handlers.TokoHandler) {
	toko := router.Group("/toko", middleware.AuthMiddleware())
	{
		toko.GET("/my", tokoHandler.GetMyToko)
		toko.GET("/:id_toko", tokoHandler.GetTokoByID)
		toko.GET("", tokoHandler.GetAllToko)
		toko.PUT("/:id_toko", tokoHandler.UpdateToko)
	}
}
