package services

import (
	"biography-api/models"
	"biography-api/repositories"
)

type ExperienceService interface {
	GetAll(page int, limit int) ([]models.Experience, int64, error)
	GetByID(id uint) (models.Experience, error)
	Create(experience models.Experience) (models.Experience, error)
	Update(id uint, experience models.Experience) (models.Experience, error)
	Delete(id uint) error
}

type experienceService struct {
	repository repositories.ExperienceRepository
}

func NewExperienceService(repository repositories.ExperienceRepository) ExperienceService {
	return &experienceService{repository}
}

func (s *experienceService) GetAll(page int, limit int) ([]models.Experience, int64, error) {
	return s.repository.GetAll(page, limit)
}

func (s *experienceService) GetByID(id uint) (models.Experience, error) {
	return s.repository.GetByID(id)
}

func (s *experienceService) Create(experience models.Experience) (models.Experience, error) {
	return s.repository.Create(experience)
}

func (s *experienceService) Update(id uint, experience models.Experience) (models.Experience, error) {
	return s.repository.Update(id, experience)
}

func (s *experienceService) Delete(id uint) error {
	return s.repository.Delete(id)
}
