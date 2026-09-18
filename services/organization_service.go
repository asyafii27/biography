package services

import (
	"biography-api/models"
	"biography-api/repositories"
)

type OrganizationService interface {
	GetAll(page int, limit int) ([]models.Organization, int64, error)
	GetByID(id uint) (models.Organization, error)
	Create(organization models.Organization) (models.Organization, error)
	Update(id uint, organization models.Organization) (models.Organization, error)
	Delete(id uint) error
}

type organizationService struct {
	repo repositories.OrganizationRepository
}

func NewOrganizationService(repo repositories.OrganizationRepository) OrganizationService {
	return &organizationService{repo}
}

func (s *organizationService) GetAll(page int, limit int) ([]models.Organization, int64, error) {
	return s.repo.GetAll(page, limit)
}

func (s *organizationService) GetByID(id uint) (models.Organization, error) {
	return s.repo.GetByID(id)
}

func (s *organizationService) Create(organization models.Organization) (models.Organization, error) {
	return s.repo.Create(organization)
}

func (s *organizationService) Update(id uint, organization models.Organization) (models.Organization, error) {
	return s.repo.Update(id, organization)
}

func (s *organizationService) Delete(id uint) error {
	return s.repo.Delete(id)
}
