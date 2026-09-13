package routes

import (
	"mini-project-pbi/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterProvCityRoutes registers EMSIFA Indonesian province and city endpoints
func RegisterProvCityRoutes(router *gin.Engine, provCityHandler *handlers.ProvCityHandler) {
	provCity := router.Group("/provcity")
	{
		provCity.GET("/listprovincies", provCityHandler.GetListProvincies)
		provCity.GET("/listcities/:prov_id", provCityHandler.GetListCities)
		provCity.GET("/detailprovince/:prov_id", provCityHandler.GetDetailProvince)
		provCity.GET("/detailcity/:city_id", provCityHandler.GetDetailCity)
	}
}
