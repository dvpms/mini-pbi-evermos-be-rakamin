package handlers

import (
	"net/http"

	"mini-project-pbi/models"
	"mini-project-pbi/services"

	"github.com/gin-gonic/gin"
)

// AuthHandler handles HTTP requests for authentication
type AuthHandler struct {
	service services.AuthService
}

// NewAuthHandler creates a new AuthHandler instance
func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to POST data", err.Error()))
		return
	}

	if err := h.service.Register(req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to POST data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to POST data", "Register Succeed"))
}

// Login handles user login and token generation
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Failed to POST data", err.Error()))
		return
	}

	data, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Failed to POST data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse("Succeed to POST data", data))
}
