package services

import (
	"go-sqlite-crud/models"
	"go-sqlite-crud/repositories"
)

type ProductService interface {
	CreateProduct(product *models.Product) error
	GetProducts(search string, categoryID int, minPrice, maxPrice float64, page, limit int) ([]models.Product, int64, error)
	GetProductByID(id uint) (models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
}

type productService struct {
	productRepo  repositories.ProductRepository
	categoryRepo repositories.CategoryRepository
}

func NewProductService(productRepo repositories.ProductRepository, categoryRepo repositories.CategoryRepository) ProductService {
	return &productService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *productService) CreateProduct(product *models.Product) error {
	// Verify Category exists
	_, err := s.categoryRepo.FindByID(product.CategoryID)
	if err != nil {
		return err
	}
	return s.productRepo.Create(product)
}

func (s *productService) GetProducts(search string, categoryID int, minPrice, maxPrice float64, page, limit int) ([]models.Product, int64, error) {
	return s.productRepo.FindAll(search, categoryID, minPrice, maxPrice, page, limit)
}

func (s *productService) GetProductByID(id uint) (models.Product, error) {
	return s.productRepo.FindByID(id)
}

func (s *productService) UpdateProduct(product *models.Product) error {
	// Verify Category exists
	_, err := s.categoryRepo.FindByID(product.CategoryID)
	if err != nil {
		return err
	}
	return s.productRepo.Update(product)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.productRepo.Delete(id)
}
