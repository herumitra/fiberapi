package models

import (
	"time"
)

type TokenBlacklist struct {
	ID        uint      `gorm:"primaryKey"`
	Token     string    `gorm:"uniqueIndex"`
	ExpiresAt time.Time `gorm:"not null"`
}
