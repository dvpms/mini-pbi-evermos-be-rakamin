package handlers

import (
	"net/http"

	"mini-project-pbi/models"
	"mini-project-pbi/services"

	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests for user profile operations
type UserHandler struct {
	service services.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetProfile returns logged-in user profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized", "User tidak terautentikasi"))
		return
	}
	userID := userIDVal.(uint)

	user, err := h.service.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse("Failed to GET data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", user))
}

// UpdateProfile updates logged-in user profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized", "User tidak terautentikasi"))
		return
	}
	userID := userIDVal.(uint)

	var req models.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to POST data", err.Error()))
		return
	}

	if err := h.service.UpdateProfile(userID, req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to POST data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to GET data", ""))
}
