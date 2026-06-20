package repositories

import (
	"go-sqlite-crud/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	FindAll() ([]models.Category, error)
	FindByID(id uint) (models.Category, error)
	Update(category *models.Category) error
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) FindByID(id uint) (models.Category, error) {
	var category models.Category
	err := r.db.Preload("Products").First(&category, id).Error
	return category, err
}

func (r *categoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	var category models.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&category).Error
}
