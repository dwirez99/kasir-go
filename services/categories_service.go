package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type CategoriesService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoriesService {
	return &CategoriesService{repo: repo}
}

func (s *CategoriesService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoriesService) Create(data *models.Category) error {
	return s.repo.Create(data)
}

func (s *CategoriesService) GetByID(id int) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoriesService) Update(category *models.Category) error {
	return s.repo.Update(category)
}

func (s *CategoriesService) Delete(id int) error {
	return s.repo.Delete(id)
}

