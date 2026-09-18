package repositories

import (
	"biography-api/models"

	"gorm.io/gorm"
)

type TechnicalExperienceRepository interface {
	GetAll(page int, limit int) ([]models.TechnicalExperience, int64, error)
	GetByID(id uint) (models.TechnicalExperience, error)
	Create(technicalExperience models.TechnicalExperience) (models.TechnicalExperience, error)
	Update(id uint, technicalExperience models.TechnicalExperience) (models.TechnicalExperience, error)
	Delete(id uint) error
}

type technicalExperienceRepository struct {
	db *gorm.DB
}

func NewTechnicalExperienceRepository(db *gorm.DB) TechnicalExperienceRepository {
	return &technicalExperienceRepository{db}
}

func (r *technicalExperienceRepository) GetAll(page int, limit int) ([]models.TechnicalExperience, int64, error) {
	var items []models.TechnicalExperience
	var totalItems int64

	r.db.Model(&models.TechnicalExperience{}).Count(&totalItems)
	offset := (page - 1) * limit
	err := r.db.Offset(offset).Limit(limit).Find(&items).Error

	return items, totalItems, err
}

func (r *technicalExperienceRepository) GetByID(id uint) (models.TechnicalExperience, error) {
	var item models.TechnicalExperience
	err := r.db.First(&item, id).Error
	return item, err
}

func (r *technicalExperienceRepository) Create(technicalExperience models.TechnicalExperience) (models.TechnicalExperience, error) {
	err := r.db.Create(&technicalExperience).Error
	return technicalExperience, err
}

func (r *technicalExperienceRepository) Update(id uint, technicalExperience models.TechnicalExperience) (models.TechnicalExperience, error) {
	var item models.TechnicalExperience
	err := r.db.First(&item, id).Error
	if err != nil {
		return item, err
	}

	err = r.db.Model(&item).Updates(technicalExperience).Error
	return item, err
}

func (r *technicalExperienceRepository) Delete(id uint) error {
	err := r.db.Delete(&models.TechnicalExperience{}, id).Error
	return err
}
