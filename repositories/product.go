package repositories

import (
	"go-sqlite-crud/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	FindAll(search string, categoryID int, minPrice, maxPrice float64, page, limit int) ([]models.Product, int64, error)
	FindByID(id uint) (models.Product, error)
	Update(product *models.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindAll(search string, categoryID int, minPrice, maxPrice float64, page, limit int) ([]models.Product, int64, error) {
	query := r.db.Model(&models.Product{}).Preload("Category")

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

func (r *productRepository) FindByID(id uint) (models.Product, error) {
	var product models.Product
	err := r.db.Preload("Category").First(&product, id).Error
	return product, err
}

func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&product).Error
}
