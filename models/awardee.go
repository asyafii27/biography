package models

import (
	"time"
)

type Awardee struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title"`
	Date        *time.Time `json:"date" gorm:"type:date"`
	Description string     `json:"description" gorm:"type:text"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
