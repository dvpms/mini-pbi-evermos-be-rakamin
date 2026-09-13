package routes

import (
	"mini-project-pbi/handlers"
	"mini-project-pbi/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterProductRoutes registers product catalog, search, and store mutation endpoints
func RegisterProductRoutes(router *gin.Engine, produkHandler *handlers.ProdukHandler) {
	// Public product routes
	product := router.Group("/product")
	{
		product.GET("", produkHandler.GetAllProduk)
		product.GET("/:id", produkHandler.GetProdukByID)
	}

	// Protected product mutation routes
	productAuth := router.Group("/product", middleware.AuthMiddleware())
	{
		productAuth.POST("", produkHandler.CreateProduk)
		productAuth.PUT("/:id", produkHandler.UpdateProduk)
		productAuth.DELETE("/:id", produkHandler.DeleteProduk)
	}
}
