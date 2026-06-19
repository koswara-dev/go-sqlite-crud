package services

import (
	"go-sqlite-crud/config"
	"go-sqlite-crud/models"
)

func CreateProduct(product *models.Product) error {
	// Verify Category exists
	var category models.Category
	if err := config.DB.First(&category, product.CategoryID).Error; err != nil {
		return err
	}
	return config.DB.Create(product).Error
}

func GetProducts(search string, categoryID int, minPrice, maxPrice float64, page, limit int) ([]models.Product, int64, error) {
	query := config.DB.Model(&models.Product{}).Preload("Category")

	if search != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	if minPrice > 0 {
		query = query.Where("price >= ?", minPrice)
	}

	if maxPrice > 0 {
		query = query.Where("price <= ?", maxPrice)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	var products []models.Product
	if err := query.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func GetProductByID(id uint) (models.Product, error) {
	var product models.Product
	err := config.DB.Preload("Category").First(&product, id).Error
	return product, err
}

func UpdateProduct(product *models.Product) error {
	// Verify Category exists
	var category models.Category
	if err := config.DB.First(&category, product.CategoryID).Error; err != nil {
		return err
	}
	return config.DB.Save(product).Error
}

func DeleteProduct(id uint) error {
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return err
	}
	return config.DB.Delete(&product).Error
}
