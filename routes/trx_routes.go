package routes

import (
	"mini-project-pbi/handlers"
	"mini-project-pbi/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterTrxRoutes registers checkout, transaction history, and detail endpoints
func RegisterTrxRoutes(router *gin.Engine, transaksiHandler *handlers.TransaksiHandler) {
	trx := router.Group("/trx", middleware.AuthMiddleware())
	{
		trx.POST("", transaksiHandler.CreateTrx)
		trx.GET("", transaksiHandler.GetAllTrx)
		trx.GET("/:id", transaksiHandler.GetTrxByID)
	}
}
