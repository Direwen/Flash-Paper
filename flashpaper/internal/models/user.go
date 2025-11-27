package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
