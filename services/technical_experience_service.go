package services

import (
	"biography-api/models"
	"biography-api/repositories"
)

type TechnicalExperienceService interface {
	GetAll(page int, limit int) ([]models.TechnicalExperience, int64, error)
	GetByID(id uint) (models.TechnicalExperience, error)
	Create(technicalExperience models.TechnicalExperience) (models.TechnicalExperience, error)
	Update(id uint, technicalExperience models.TechnicalExperience) (models.TechnicalExperience, error)
	Delete(id uint) error
}

type technicalExperienceService struct {
	repo repositories.TechnicalExperienceRepository
}

func NewTechnicalExperienceService(repo repositories.TechnicalExperienceRepository) TechnicalExperienceService {
	return &technicalExperienceService{repo}
}

func (s *technicalExperienceService) GetAll(page int, limit int) ([]models.TechnicalExperience, int64, error) {
	return s.repo.GetAll(page, limit)
}

func (s *technicalExperienceService) GetByID(id uint) (models.TechnicalExperience, error) {
	return s.repo.GetByID(id)
}

func (s *technicalExperienceService) Create(technicalExperience models.TechnicalExperience) (models.TechnicalExperience, error) {
	return s.repo.Create(technicalExperience)
}

func (s *technicalExperienceService) Update(id uint, technicalExperience models.TechnicalExperience) (models.TechnicalExperience, error) {
	return s.repo.Update(id, technicalExperience)
}

func (s *technicalExperienceService) Delete(id uint) error {
	return s.repo.Delete(id)
}
