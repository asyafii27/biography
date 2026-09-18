package services

import (
	"biography-api/models"
	"biography-api/repositories"
)

type AwardeeService interface {
	GetAll(page int, limit int) ([]models.Awardee, int64, error)
	GetByID(id uint) (models.Awardee, error)
	Create(awardee models.Awardee) (models.Awardee, error)
	Update(id uint, awardee models.Awardee) (models.Awardee, error)
	Delete(id uint) error
}

type awardeeService struct {
	repo repositories.AwardeeRepository
}

func NewAwardeeService(repo repositories.AwardeeRepository) AwardeeService {
	return &awardeeService{repo}
}

func (s *awardeeService) GetAll(page int, limit int) ([]models.Awardee, int64, error) {
	return s.repo.GetAll(page, limit)
}

func (s *awardeeService) GetByID(id uint) (models.Awardee, error) {
	return s.repo.GetByID(id)
}

func (s *awardeeService) Create(awardee models.Awardee) (models.Awardee, error) {
	return s.repo.Create(awardee)
}

func (s *awardeeService) Update(id uint, awardee models.Awardee) (models.Awardee, error) {
	return s.repo.Update(id, awardee)
}

func (s *awardeeService) Delete(id uint) error {
	return s.repo.Delete(id)
}
