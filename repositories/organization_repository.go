package repositories

import (
	"biography-api/models"

	"gorm.io/gorm"
)

type OrganizationRepository interface {
	GetAll(page int, limit int) ([]models.Organization, int64, error)
	GetByID(id uint) (models.Organization, error)
	Create(organization models.Organization) (models.Organization, error)
	Update(id uint, organization models.Organization) (models.Organization, error)
	Delete(id uint) error
}

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{db}
}

func (r *organizationRepository) GetAll(page int, limit int) ([]models.Organization, int64, error) {
	var items []models.Organization
	var totalItems int64

	r.db.Model(&models.Organization{}).Count(&totalItems)
	offset := (page - 1) * limit
	err := r.db.Offset(offset).Limit(limit).Find(&items).Error

	return items, totalItems, err
}

func (r *organizationRepository) GetByID(id uint) (models.Organization, error) {
	var item models.Organization
	err := r.db.First(&item, id).Error
	return item, err
}

func (r *organizationRepository) Create(organization models.Organization) (models.Organization, error) {
	err := r.db.Create(&organization).Error
	return organization, err
}

func (r *organizationRepository) Update(id uint, organization models.Organization) (models.Organization, error) {
	var item models.Organization
	err := r.db.First(&item, id).Error
	if err != nil {
		return item, err
	}

	err = r.db.Model(&item).Updates(organization).Error
	return item, err
}

func (r *organizationRepository) Delete(id uint) error {
	err := r.db.Delete(&models.Organization{}, id).Error
	return err
}
