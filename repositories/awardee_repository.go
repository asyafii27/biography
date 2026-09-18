package repositories

import (
	"biography-api/models"

	"gorm.io/gorm"
)

type AwardeeRepository interface {
	GetAll(page int, limit int) ([]models.Awardee, int64, error)
	GetByID(id uint) (models.Awardee, error)
	Create(awardee models.Awardee) (models.Awardee, error)
	Update(id uint, awardee models.Awardee) (models.Awardee, error)
	Delete(id uint) error
}

type awardeeRepository struct {
	db *gorm.DB
}

func NewAwardeeRepository(db *gorm.DB) AwardeeRepository {
	return &awardeeRepository{db}
}

func (r *awardeeRepository) GetAll(page int, limit int) ([]models.Awardee, int64, error) {
	var items []models.Awardee
	var totalItems int64

	r.db.Model(&models.Awardee{}).Count(&totalItems)
	offset := (page - 1) * limit
	err := r.db.Offset(offset).Limit(limit).Find(&items).Error

	return items, totalItems, err
}

func (r *awardeeRepository) GetByID(id uint) (models.Awardee, error) {
	var item models.Awardee
	err := r.db.First(&item, id).Error
	return item, err
}

func (r *awardeeRepository) Create(awardee models.Awardee) (models.Awardee, error) {
	err := r.db.Create(&awardee).Error
	return awardee, err
}

func (r *awardeeRepository) Update(id uint, awardee models.Awardee) (models.Awardee, error) {
	var item models.Awardee
	err := r.db.First(&item, id).Error
	if err != nil {
		return item, err
	}

	err = r.db.Model(&item).Updates(awardee).Error
	return item, err
}

func (r *awardeeRepository) Delete(id uint) error {
	err := r.db.Delete(&models.Awardee{}, id).Error
	return err
}
