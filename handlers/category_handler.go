package handlers

import (
	"net/http"
	"strconv"

	"mini-project-pbi/models"
	"mini-project-pbi/services"

	"github.com/gin-gonic/gin"
)

// CategoryHandler handles HTTP requests for categories
type CategoryHandler struct {
	categoryService services.CategoryService
}

// NewCategoryHandler creates a new CategoryHandler instance
func NewCategoryHandler(categoryService services.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

// GetAllCategories returns list of all categories
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.categoryService.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", categories))
}

// GetCategoryByID returns single category by ID
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to GET data", "Invalid category ID"))
		return
	}

	category, err := h.categoryService.GetCategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", category))
}

// CreateCategory creates a new category (Admin only)
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req models.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to POST data", err.Error()))
		return
	}

	id, err := h.categoryService.CreateCategory(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse("Failed to POST data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to POST data", id))
}

// UpdateCategory updates existing category (Admin only)
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to PUT data", "Invalid category ID"))
		return
	}

	var req models.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to PUT data", err.Error()))
		return
	}

	if err := h.categoryService.UpdateCategory(uint(id), req); err != nil {
		if err.Error() == "Kategori tidak ditemukan" {
			c.JSON(http.StatusNotFound, models.ErrorResponse("Failed to PUT data", err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse("Failed to PUT data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", ""))
}

// DeleteCategory deletes a category by ID (Admin only)
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to DELETE data", "Invalid category ID"))
		return
	}

	if err := h.categoryService.DeleteCategory(uint(id)); err != nil {
		if err.Error() == "Kategori tidak ditemukan" {
			c.JSON(http.StatusNotFound, models.ErrorResponse("Failed to DELETE data", err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse("Failed to DELETE data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", ""))
}
