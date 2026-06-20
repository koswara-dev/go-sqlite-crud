package services

import (
	"go-sqlite-crud/models"
	"go-sqlite-crud/repositories"
)

type CategoryService interface {
	CreateCategory(category *models.Category) error
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id uint) (models.Category, error)
	UpdateCategory(category *models.Category) error
	DeleteCategory(id uint) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(category *models.Category) error {
	return s.repo.Create(category)
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) GetCategoryByID(id uint) (models.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) UpdateCategory(category *models.Category) error {
	return s.repo.Update(category)
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.repo.Delete(id)
}
