package services

import (
	"strings"
	"time"

	"github.com/direwen/flashpaper/internal/models"
	"github.com/direwen/flashpaper/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SnippetService struct {
	db *gorm.DB
}

func NewSnippetService(db *gorm.DB) *SnippetService {
	return &SnippetService{db: db}
}

func (s SnippetService) CreateSnippet(
	userID *uuid.UUID,
	content,
	title,
	language string,
	maxViews int,
	expiresInMinutes int,
) (*models.Snippet, error) {

	// Encrypt Content
	content = strings.TrimSpace(content)
	encrypted, err := utils.Encrypt(content)
	if err != nil {
		return nil, err
	}

	// Sanitize Title
	title = strings.TrimSpace(title)
	// Sanitize language
	language = strings.ToLower(strings.TrimSpace(language))

	// Calc Expiry
	expiresAt := time.Now().Add(time.Minute * time.Duration(expiresInMinutes))

	// Prepare Model
	snippet := &models.Snippet{
		UserID:    userID,
		Content:   encrypted,
		Title:     title,
		Language:  utils.SanitizeLanguage(language),
		MaxViews:  maxViews,
		ExpiresAt: expiresAt,
	}

	if err := s.db.Create(snippet).Error; err != nil {
		return nil, err
	}

	return snippet, nil
}
