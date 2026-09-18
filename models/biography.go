package models

import (
	"time"
)

type Biography struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	Name            string     `json:"name" gorm:"not null"`
	Nickname        string     `json:"nickname" gorm:"not null"`
	BirthDate       *time.Time `json:"birth_date" gorm:"type:date;not null"`
	Gender          string     `json:"gender" gorm:"not null"`
	FullAddress     string     `json:"full_address" gorm:"type:text;not null"`
	MobileNumber    string     `json:"mobile_number" gorm:"not null"`
	JobTitle        string     `json:"job_title" gorm:"not null"`
	ProfilePhotoURL *string    `json:"profile_photo_url"`
	Motto           *string    `json:"motto"`
	Description          string                `json:"description" gorm:"type:text;not null"`
	CreatedAt            time.Time             `json:"created_at"`
	UpdatedAt            time.Time             `json:"updated_at"`

	Experiences          []Experience          `json:"experiences" gorm:"foreignKey:BiographyID"`
	Awardees             []Awardee             `json:"awardees" gorm:"foreignKey:BiographyID"`
	Organizations        []Organization        `json:"organizations" gorm:"foreignKey:BiographyID"`
	Skills               []Skill               `json:"skills" gorm:"foreignKey:BiographyID"`
	TechnicalExperiences []TechnicalExperience `json:"technical_experiences" gorm:"foreignKey:BiographyID"`
}
