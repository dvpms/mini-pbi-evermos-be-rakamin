package handlers

import (
	"net/http"
	"strconv"

	"mini-project-pbi/models"
	"mini-project-pbi/services"

	"github.com/gin-gonic/gin"
)

// AlamatHandler handles HTTP requests for shipping addresses
type AlamatHandler struct {
	service services.AlamatService
}

// NewAlamatHandler creates a new AlamatHandler instance
func NewAlamatHandler(service services.AlamatService) *AlamatHandler {
	return &AlamatHandler{service: service}
}

// GetMyAlamat returns all addresses of the logged-in user
func (h *AlamatHandler) GetMyAlamat(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized", "User tidak terautentikasi"))
		return
	}
	userID := userIDVal.(uint)

	alamats, err := h.service.GetAlamatByUser(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}

	if alamats == nil {
		alamats = []models.Alamat{}
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", alamats))
}

// GetAlamatByID returns a specific address by ID
func (h *AlamatHandler) GetAlamatByID(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized", "User tidak terautentikasi"))
		return
	}
	userID := userIDVal.(uint)

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", "Invalid ID"))
		return
	}

	alamat, err := h.service.GetAlamatByID(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", alamat))
}

// CreateAlamat adds a new address for logged-in user
func (h *AlamatHandler) CreateAlamat(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized", "User tidak terautentikasi"))
		return
	}
	userID := userIDVal.(uint)

	var req models.CreateAlamatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to POST data", err.Error()))
		return
	}

	id, err := h.service.CreateAlamat(userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to POST data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to POST data", id))
}

// UpdateAlamat updates an existing address
func (h *AlamatHandler) UpdateAlamat(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized", "User tidak terautentikasi"))
		return
	}
	userID := userIDVal.(uint)

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", "Invalid ID"))
		return
	}

	var req models.UpdateAlamatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}

	if err := h.service.UpdateAlamat(uint(id), userID, req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", ""))
}

// DeleteAlamat deletes an address
func (h *AlamatHandler) DeleteAlamat(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized", "User tidak terautentikasi"))
		return
	}
	userID := userIDVal.(uint)

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", "Invalid ID"))
		return
	}

	if err := h.service.DeleteAlamat(uint(id), userID); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", ""))
}
