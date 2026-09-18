package repositories

import (
	"biography-api/models"

	"gorm.io/gorm"
)

type SkillRepository interface {
	GetAll(page int, limit int) ([]models.Skill, int64, error)
	GetByID(id uint) (models.Skill, error)
	Create(skill models.Skill) (models.Skill, error)
	Update(id uint, skill models.Skill) (models.Skill, error)
	Delete(id uint) error
}

type skillRepository struct {
	db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) SkillRepository {
	return &skillRepository{db}
}

func (r *skillRepository) GetAll(page int, limit int) ([]models.Skill, int64, error) {
	var items []models.Skill
	var totalItems int64

	r.db.Model(&models.Skill{}).Count(&totalItems)
	offset := (page - 1) * limit
	err := r.db.Offset(offset).Limit(limit).Find(&items).Error

	return items, totalItems, err
}

func (r *skillRepository) GetByID(id uint) (models.Skill, error) {
	var item models.Skill
	err := r.db.First(&item, id).Error
	return item, err
}

func (r *skillRepository) Create(skill models.Skill) (models.Skill, error) {
	err := r.db.Create(&skill).Error
	return skill, err
}

func (r *skillRepository) Update(id uint, skill models.Skill) (models.Skill, error) {
	var item models.Skill
	err := r.db.First(&item, id).Error
	if err != nil {
		return item, err
	}

	err = r.db.Model(&item).Updates(skill).Error
	return item, err
}

func (r *skillRepository) Delete(id uint) error {
	err := r.db.Delete(&models.Skill{}, id).Error
	return err
}
