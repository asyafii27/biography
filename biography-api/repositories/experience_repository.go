package repositories

import (
	"biography-api/models"

	"gorm.io/gorm"
)

type ExperienceRepository interface {
	GetAll(page int, limit int) ([]models.Experience, int64, error)
	GetByID(id uint) (models.Experience, error)
	Create(experience models.Experience) (models.Experience, error)
	Update(id uint, experience models.Experience) (models.Experience, error)
	Delete(id uint) error
}

type experienceRepository struct {
	db *gorm.DB
}

func NewExperienceRepository(db *gorm.DB) ExperienceRepository {
	return &experienceRepository{db}
}

func (r *experienceRepository) GetAll(page int, limit int) ([]models.Experience, int64, error) {
	var experiences []models.Experience
	var totalItems int64

	// Count total records
	r.db.Model(&models.Experience{}).Count(&totalItems)

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch data with limit and offset
	err := r.db.Offset(offset).Limit(limit).Find(&experiences).Error

	return experiences, totalItems, err
}

func (r *experienceRepository) GetByID(id uint) (models.Experience, error) {
	var experience models.Experience
	err := r.db.First(&experience, id).Error
	return experience, err
}

func (r *experienceRepository) Create(experience models.Experience) (models.Experience, error) {
	err := r.db.Create(&experience).Error
	return experience, err
}

func (r *experienceRepository) Update(id uint, experience models.Experience) (models.Experience, error) {
	var exp models.Experience
	err := r.db.First(&exp, id).Error
	if err != nil {
		return exp, err
	}

	err = r.db.Model(&exp).Updates(experience).Error
	return exp, err
}

func (r *experienceRepository) Delete(id uint) error {
	err := r.db.Delete(&models.Experience{}, id).Error
	return err
}
