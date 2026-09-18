package services

import (
	"biography-api/models"
	"biography-api/repositories"
)

type SkillService interface {
	GetAll(page int, limit int) ([]models.Skill, int64, error)
	GetByID(id uint) (models.Skill, error)
	Create(skill models.Skill) (models.Skill, error)
	Update(id uint, skill models.Skill) (models.Skill, error)
	Delete(id uint) error
}

type skillService struct {
	repo repositories.SkillRepository
}

func NewSkillService(repo repositories.SkillRepository) SkillService {
	return &skillService{repo}
}

func (s *skillService) GetAll(page int, limit int) ([]models.Skill, int64, error) {
	return s.repo.GetAll(page, limit)
}

func (s *skillService) GetByID(id uint) (models.Skill, error) {
	return s.repo.GetByID(id)
}

func (s *skillService) Create(skill models.Skill) (models.Skill, error) {
	return s.repo.Create(skill)
}

func (s *skillService) Update(id uint, skill models.Skill) (models.Skill, error) {
	return s.repo.Update(id, skill)
}

func (s *skillService) Delete(id uint) error {
	return s.repo.Delete(id)
}
