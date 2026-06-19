package services

import (
	"go-sqlite-crud/config"
	"go-sqlite-crud/models"
)

func CreateCategory(category *models.Category) error {
	return config.DB.Create(category).Error
}

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := config.DB.Find(&categories).Error
	return categories, err
}

func GetCategoryByID(id uint) (models.Category, error) {
	var category models.Category
	err := config.DB.Preload("Products").First(&category, id).Error
	return category, err
}

func UpdateCategory(category *models.Category) error {
	return config.DB.Save(category).Error
}

func DeleteCategory(id uint) error {
	// First retrieve category to make sure it exists and to comply with GORM delete triggers if any
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return err
	}
	return config.DB.Delete(&category).Error
}
