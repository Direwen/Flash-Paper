package services

import (
	"errors"
	"strings"
	"time"

	"github.com/direwen/flashpaper/internal/models"
	"github.com/direwen/flashpaper/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (s SnippetService) GetSnippet(snippetID string) (*models.Snippet, error) {
	var snippet models.Snippet

	// Validate uuid format
	uid, err := uuid.Parse(snippetID)
	if err != nil {
		return nil, err
	}

	// Start Transaction
	tx := s.db.Begin()

	// If anything panics, Rollback changes
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Lock the row
	if err := tx.Clauses(clause.Locking{
		Strength: clause.LockingStrengthUpdate,
	}).First(&snippet, uid).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Expired?
	if time.Now().After(snippet.ExpiresAt) {
		tx.Rollback()
		return nil, errors.New("expired")
	}

	// Burnt? (Views > MaxViews)
	if snippet.CurrentViews >= snippet.MaxViews {
		tx.Rollback()
		return nil, errors.New("burnt")
	}

	// Increment View
	snippet.CurrentViews++
	if err := tx.Save(&snippet).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// If successful, decrypt the content
	decrypted, err := utils.Decrypt(snippet.Content)
	if err != nil {
		return nil, errors.New("decryption_failed")
	}
	snippet.Content = decrypted

	return &snippet, nil
}

func (s SnippetService) DeleteSnippet(snippetID uuid.UUID, userID uuid.UUID) error {
	result := s.db.Where("id = ? AND user_id = ?", snippetID, userID).Delete(&models.Snippet{})
	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return errors.New("not_found")
	}

	return nil
}
