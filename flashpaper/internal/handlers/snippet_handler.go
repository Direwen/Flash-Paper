package handlers

import (
	"net/http"

	"github.com/direwen/flashpaper/internal/services"
	"github.com/direwen/flashpaper/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SnippetHandler struct {
	service *services.SnippetService
}

func NewSnippetHandler(service *services.SnippetService) *SnippetHandler {
	return &SnippetHandler{service: service}
}

type CreateSnippetRequest struct {
	Content   string `json:"content" binding:"required"`
	Title     string `json:"title"`
	Language  string `json:"language"`
	MaxViews  int    `json:"max_views" binding:"required,min=1"`
	ExpiresIn int    `json:"expires_in" binding:"required,min=1"`
}

func (h *SnippetHandler) Create(c *gin.Context) {
	var req CreateSnippetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	// Get User ID from the context (middleware assigned the user id)
	userIDVal, exists := c.Get("userID")
	var userID *uuid.UUID
	if exists {
		id := userIDVal.(uuid.UUID)
		userID = &id
	}

	snippet, err := h.service.CreateSnippet(
		userID,
		req.Content,
		req.Title,
		req.Language,
		req.MaxViews,
		req.ExpiresIn,
	)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	fullLink := "/api/snippets/" + snippet.ID.String()

	utils.SendSuccess(c, http.StatusCreated, gin.H{
		"message":    "Snippet created successfully",
		"id":         snippet.ID,
		"link":       fullLink,
		"expires_at": snippet.ExpiresAt,
		"max_views":  snippet.MaxViews,
	})
}
