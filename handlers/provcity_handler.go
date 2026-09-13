package handlers

import (
	"net/http"

	"mini-project-pbi/models"
	"mini-project-pbi/services"

	"github.com/gin-gonic/gin"
)

// ProvCityHandler handles HTTP requests for Indonesian administrative regions
type ProvCityHandler struct {
	service services.ProvCityService
}

// NewProvCityHandler creates a new ProvCityHandler instance
func NewProvCityHandler(service services.ProvCityService) *ProvCityHandler {
	return &ProvCityHandler{service: service}
}

// GetListProvincies returns list of all provinces in Indonesia
func (h *ProvCityHandler) GetListProvincies(c *gin.Context) {
	data, err := h.service.GetProvinces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to get data", data))
}

// GetListCities returns list of cities in a province
func (h *ProvCityHandler) GetListCities(c *gin.Context) {
	provID := c.Param("prov_id")
	data, err := h.service.GetCities(provID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to get data", data))
}

// GetDetailProvince returns details of a single province
func (h *ProvCityHandler) GetDetailProvince(c *gin.Context) {
	provID := c.Param("prov_id")
	data, err := h.service.GetDetailProvince(provID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to get data", data))
}

// GetDetailCity returns details of a single city
func (h *ProvCityHandler) GetDetailCity(c *gin.Context) {
	cityID := c.Param("city_id")
	data, err := h.service.GetDetailCity(cityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to get data", data))
}
