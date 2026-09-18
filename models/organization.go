package models

import (
	"time"
)

type Organization struct {
	ID               uint       `json:"id" gorm:"primaryKey"`
	Role             string     `json:"role"`
	OrganizationName string     `json:"organization_name"`
	StartDate        *time.Time `json:"start_date" gorm:"type:date"`
	EndDate          *time.Time `json:"end_date" gorm:"type:date"`
	IsCurrent        bool       `json:"is_current"`
	Description      string     `json:"description" gorm:"type:text"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}
