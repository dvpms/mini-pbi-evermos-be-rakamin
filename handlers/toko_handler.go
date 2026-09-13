package handlers

import (
	"net/http"
	"strconv"

	"mini-project-pbi/models"
	"mini-project-pbi/services"

	"github.com/gin-gonic/gin"
)

// TokoHandler handles HTTP requests for store management
type TokoHandler struct {
	tokoService services.TokoService
}

// NewTokoHandler creates a new TokoHandler instance
func NewTokoHandler(tokoService services.TokoService) *TokoHandler {
	return &TokoHandler{tokoService: tokoService}
}

// GetMyToko retrieves the authenticated user's store
func (h *TokoHandler) GetMyToko(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Failed to GET data", "Unauthorized"))
		return
	}
	userID := userIDVal.(uint)

	toko, err := h.tokoService.GetMyToko(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", toko))
}

// GetTokoByID retrieves public store detail by ID
func (h *TokoHandler) GetTokoByID(c *gin.Context) {
	idParam := c.Param("id_toko")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", "Invalid ID toko parameter"))
		return
	}

	toko, err := h.tokoService.GetTokoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse("Failed to GET data", "Toko tidak ditemukan"))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", toko))
}

// GetAllToko retrieves paginated list of stores with optional search filter
func (h *TokoHandler) GetAllToko(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	nama := c.Query("nama")

	results, _, err := h.tokoService.GetAllToko(page, limit, nama)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}

	dataResponse := gin.H{
		"page":  page,
		"limit": limit,
		"data":  results,
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", dataResponse))
}

// UpdateToko updates store information and photo
func (h *TokoHandler) UpdateToko(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Failed to UPDATE data", "Unauthorized"))
		return
	}
	userID := userIDVal.(uint)

	idParam := c.Param("id_toko")
	tokoID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to UPDATE data", "Invalid ID toko parameter"))
		return
	}

	namaToko := c.PostForm("nama_toko")
	file, _ := c.FormFile("photo")

	if err := h.tokoService.UpdateToko(uint(tokoID), userID, namaToko, file); err != nil {
		if err.Error() == "Anda tidak memiliki akses untuk mengubah toko ini" {
			c.JSON(http.StatusForbidden, models.ErrorResponse("Failed to UPDATE data", err.Error()))
			return
		}
		if err.Error() == "Toko tidak ditemukan" {
			c.JSON(http.StatusNotFound, models.ErrorResponse("Failed to UPDATE data", err.Error()))
			return
		}
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to UPDATE data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to UPDATE data", "Update toko succeed"))
}
