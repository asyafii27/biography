package models

import (
	"time"
)

type TechnicalExperience struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	BiographyID uint      `json:"biography_id" gorm:"index;not null"`
	Title       string    `json:"title"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
