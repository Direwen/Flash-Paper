package models

import (
	"time"

	"github.com/google/uuid"
)

type Snippet struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key;"`
	UserID       uuid.UUID `gorm:"type:uuid;index"`
	User         User      `gorm:"foreignKey:UserID"` // Virtual field (populated only when preloaded, not stored in DB)
	Content      string    `gorm:"not null"`
	Title        string
	Language     string
	CurrentViews int       `gorm:"default:0"`
	MaxViews     int       `gorm:"default:0"`
	ExpiresAt    time.Time `gorm:"index"`
	CreatedAt    time.Time
}
