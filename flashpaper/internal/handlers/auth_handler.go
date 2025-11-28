package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/direwen/flashpaper/internal/services"
	"github.com/direwen/flashpaper/pkg/utils"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest

	// Parse & Validate JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	// Call the service
	err := h.service.RegisterUser(req.Email, req.Password)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendMessage(c, http.StatusCreated, "User registered successfully")

}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	// Parse & Validate JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	// Login
	token, err := h.service.LoginUser(req.Email, req.Password)
	if err != nil {
		utils.SendError(c, http.StatusUnauthorized, err)
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H{
		"token": token,
	})
}
