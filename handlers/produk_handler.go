package handlers

import (
	"mime/multipart"
	"net/http"
	"strconv"

	"mini-project-pbi/models"
	"mini-project-pbi/repositories"
	"mini-project-pbi/services"

	"github.com/gin-gonic/gin"
)

// ProdukHandler handles HTTP requests for product catalog
type ProdukHandler struct {
	produkService services.ProdukService
}

// NewProdukHandler creates a new ProdukHandler instance
func NewProdukHandler(produkService services.ProdukService) *ProdukHandler {
	return &ProdukHandler{produkService: produkService}
}

// GetAllProduk handles filtered and paginated product retrieval
func (h *ProdukHandler) GetAllProduk(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	namaProduk := c.Query("nama_produk")
	categoryID, _ := strconv.ParseUint(c.Query("category_id"), 10, 32)
	tokoID, _ := strconv.ParseUint(c.Query("toko_id"), 10, 32)
	minHarga, _ := strconv.ParseFloat(c.Query("min_harga"), 64)
	maxHarga, _ := strconv.ParseFloat(c.Query("max_harga"), 64)

	filter := repositories.ProdukFilter{
		NamaProduk: namaProduk,
		CategoryID: uint(categoryID),
		TokoID:     uint(tokoID),
		MinHarga:   minHarga,
		MaxHarga:   maxHarga,
		Page:       page,
		Limit:      limit,
	}

	produks, _, err := h.produkService.GetAllProduk(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}

	if produks == nil {
		produks = []models.Produk{}
	}

	dataResponse := gin.H{
		"data":  produks,
		"page":  page,
		"limit": limit,
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", dataResponse))
}

// GetProdukByID handles single product detail retrieval
func (h *ProdukHandler) GetProdukByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", "Invalid product ID"))
		return
	}

	produk, err := h.produkService.GetProdukByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse("Failed to GET data", "No Data Product"))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", produk))
}

// CreateProduk handles new product creation with photos
func (h *ProdukHandler) CreateProduk(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Failed to POST data", "Unauthorized"))
		return
	}
	userID := userIDVal.(uint)

	var req models.CreateProdukRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to POST data", err.Error()))
		return
	}

	var photos []*multipart.FileHeader
	if form, err := c.MultipartForm(); err == nil && form != nil {
		photos = form.File["photos"]
	}

	id, err := h.produkService.CreateProduk(userID, req, photos)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to POST data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to POST data", id))
}

// UpdateProduk handles updating product data and photos
func (h *ProdukHandler) UpdateProduk(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Failed to UPDATE data", "Unauthorized"))
		return
	}
	userID := userIDVal.(uint)

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", "Invalid product ID"))
		return
	}

	var req models.UpdateProdukRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}

	var photos []*multipart.FileHeader
	if form, err := c.MultipartForm(); err == nil && form != nil {
		photos = form.File["photos"]
	}

	if err := h.produkService.UpdateProduk(uint(id), userID, req, photos); err != nil {
		if err.Error() == "Anda tidak memiliki akses untuk mengubah produk ini" {
			c.JSON(http.StatusForbidden, models.ErrorResponse("Failed to GET data", err.Error()))
			return
		}
		if err.Error() == "No Data Product" || err.Error() == "Toko tidak ditemukan" {
			c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", "record not found"))
			return
		}
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", ""))
}

// DeleteProduk handles product deletion with ownership check
func (h *ProdukHandler) DeleteProduk(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Failed to GET data", "Unauthorized"))
		return
	}
	userID := userIDVal.(uint)

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", "Invalid product ID"))
		return
	}

	if err := h.produkService.DeleteProduk(uint(id), userID); err != nil {
		if err.Error() == "Anda tidak memiliki akses untuk menghapus produk ini" {
			c.JSON(http.StatusForbidden, models.ErrorResponse("Failed to GET data", err.Error()))
			return
		}
		if err.Error() == "No Data Product" || err.Error() == "Toko tidak ditemukan" {
			c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", "record not found"))
			return
		}
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", ""))
}
