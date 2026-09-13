package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	Name            string     `json:"name" gorm:"not null"`
	Email           string     `json:"email" gorm:"unique;not null"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Password        string     `json:"-" gorm:"not null"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
