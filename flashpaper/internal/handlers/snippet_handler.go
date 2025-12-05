package handlers

import (
	"errors"
	"math"
	"net/http"
	"strconv"

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
		c.Request.Context(),
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

func (h *SnippetHandler) Get(c *gin.Context) {
	// Get ID from route param
	snippetID := c.Param("id")

	snippet, err := h.service.GetSnippet(c.Request.Context(), snippetID)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, errors.New("snippet's unavailable"))
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H{
		"title":      snippet.Title,
		"content":    snippet.Content,
		"language":   snippet.Language,
		"views_left": snippet.MaxViews - snippet.CurrentViews,
		"expires_at": snippet.ExpiresAt,
		"created_at": snippet.CreatedAt,
	})
}

func (h *SnippetHandler) Delete(c *gin.Context) {
	snippetIDval := c.Param("id")
	snippetID, err := uuid.Parse(snippetIDval)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	userIDval, _ := c.Get("userID")
	userID := userIDval.(uuid.UUID)

	err = h.service.DeleteSnippet(c.Request.Context(), snippetID, userID)
	if err != nil {
		if err.Error() == "not_found" {
			utils.SendError(c, http.StatusNotFound, errors.New("snippet not found or access denied"))
		} else {
			utils.SendError(c, http.StatusInternalServerError, err)
		}
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H{
		"message": "Snippet deleted successfully",
	})
}

func (h *SnippetHandler) GetDashboard(c *gin.Context) {
	userIDval, exists := c.Get("userID")
	if !exists {
		utils.SendError(c, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	userID := userIDval.(uuid.UUID)

	stats, err := h.service.GetDashboardStats(c.Request.Context(), userID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendSuccess(c, http.StatusOK, stats)
}

func (h *SnippetHandler) List(c *gin.Context) {

	userIDVal, _ := c.Get("userID")
	userID := userIDVal.(uuid.UUID)
	// int() only works for numeric types
	// to convert string to int, strconv.Atoi is required
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	snippets, total, err := h.service.GetActiveSnippets(c.Request.Context(), userID, page, limit)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, errors.New("failed to fetch snippets"))
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H{
		"data": snippets,
		"meta": gin.H{
			"current_page": page,
			"per_page":     limit,
			"total_items":  total,
			"total_pages":  int(math.Ceil(float64(total) / float64(limit))),
		},
	})
}
