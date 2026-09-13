package handlers

import (
	"net/http"
	"strconv"

	"mini-project-pbi/models"
	"mini-project-pbi/services"

	"github.com/gin-gonic/gin"
)

// TransaksiHandler handles HTTP requests for transactions and orders
type TransaksiHandler struct {
	transaksiService services.TransaksiService
}

// NewTransaksiHandler creates a new TransaksiHandler instance
func NewTransaksiHandler(transaksiService services.TransaksiService) *TransaksiHandler {
	return &TransaksiHandler{transaksiService: transaksiService}
}

// CreateTrx handles checkout and transaction creation
func (h *TransaksiHandler) CreateTrx(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Failed to POST data", "Unauthorized"))
		return
	}
	userID := userIDVal.(uint)

	var req models.CreateTransaksiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to POST data", err.Error()))
		return
	}

	id, err := h.transaksiService.CreateTrx(userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to POST data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to POST data", id))
}

// GetAllTrx handles listing transactions for authenticated user
func (h *TransaksiHandler) GetAllTrx(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Failed to GET data", "Unauthorized"))
		return
	}
	userID := userIDVal.(uint)

	transaksis, err := h.transaksiService.GetAllTrx(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", transaksis))
}

// GetTrxByID handles retrieving transaction detail by ID for authenticated user
func (h *TransaksiHandler) GetTrxByID(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Failed to GET data", "Unauthorized"))
		return
	}
	userID := userIDVal.(uint)

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", "Invalid transaction ID"))
		return
	}

	trx, err := h.transaksiService.GetTrxByID(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse("Failed to GET data", "No Data Trx"))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", trx))
}
