package services

import (
	"biography-api/models"
	"biography-api/repositories"
)

type BiographyService interface {
	GetAll(page int, limit int) ([]models.Biography, int64, error)
	GetByID(id uint) (*models.Biography, error)
	Create(biography models.Biography) (*models.Biography, error)
	Update(id uint, biography models.Biography) (*models.Biography, error)
	Delete(id uint) error
}

type biographyService struct {
	repo repositories.BiographyRepository
}

func NewBiographyService(repo repositories.BiographyRepository) BiographyService {
	return &biographyService{repo}
}

func (s *biographyService) GetAll(page int, limit int) ([]models.Biography, int64, error) {
	return s.repo.GetAll(page, limit)
}

func (s *biographyService) GetByID(id uint) (*models.Biography, error) {
	return s.repo.GetByID(id)
}

func (s *biographyService) Create(biography models.Biography) (*models.Biography, error) {
	return s.repo.Create(biography)
}

func (s *biographyService) Update(id uint, biography models.Biography) (*models.Biography, error) {
	return s.repo.Update(id, biography)
}

func (s *biographyService) Delete(id uint) error {
	return s.repo.Delete(id)
}
