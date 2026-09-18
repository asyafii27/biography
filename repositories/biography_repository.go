package repositories

import (
	"biography-api/models"
	"gorm.io/gorm"
)

type BiographyRepository interface {
	GetAll(page int, limit int) ([]models.Biography, int64, error)
	GetByID(id uint) (*models.Biography, error)
	Create(biography models.Biography) (*models.Biography, error)
	Update(id uint, biography models.Biography) (*models.Biography, error)
	Delete(id uint) error
}

type biographyRepository struct {
	db *gorm.DB
}

func NewBiographyRepository(db *gorm.DB) BiographyRepository {
	return &biographyRepository{db}
}

func (r *biographyRepository) GetAll(page int, limit int) ([]models.Biography, int64, error) {
	var biographies []models.Biography
	var totalItems int64

	offset := (page - 1) * limit

	r.db.Model(&models.Biography{}).Count(&totalItems)
	err := r.db.Preload("Experiences").Preload("Awardees").Preload("Organizations").Preload("Skills").Preload("TechnicalExperiences").Offset(offset).Limit(limit).Find(&biographies).Error
	return biographies, totalItems, err
}

func (r *biographyRepository) GetByID(id uint) (*models.Biography, error) {
	var biography models.Biography
	err := r.db.Preload("Experiences").Preload("Awardees").Preload("Organizations").Preload("Skills").Preload("TechnicalExperiences").First(&biography, id).Error
	if err != nil {
		return nil, err
	}
	return &biography, nil
}

func (r *biographyRepository) Create(biography models.Biography) (*models.Biography, error) {
	err := r.db.Create(&biography).Error
	return &biography, err
}

func (r *biographyRepository) Update(id uint, biography models.Biography) (*models.Biography, error) {
	var existingBiography models.Biography
	err := r.db.First(&existingBiography, id).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Model(&existingBiography).Updates(biography).Error
	return &existingBiography, err
}

func (r *biographyRepository) Delete(id uint) error {
	return r.db.Delete(&models.Biography{}, id).Error
}
