package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"go-sqlite-crud/models"
	"go-sqlite-crud/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// CreateProduct handles POST /api/v1/products
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create product: " + err.Error()})
		return
	}

	// Fetch product with preloaded category for response
	resProduct, err := h.service.GetProductByID(product.ID)
	if err == nil {
		product = resProduct
	}

	c.JSON(http.StatusCreated, product)
}

// GetProducts handles GET /api/v1/products (search, filter, pagination)
func (h *ProductHandler) GetProducts(c *gin.Context) {
	search := c.Query("q")
	if search == "" {
		search = c.Query("search")
	}

	var categoryID int
	if catIDStr := c.Query("category_id"); catIDStr != "" {
		categoryID, _ = strconv.Atoi(catIDStr)
	}

	var minPrice float64
	if minPriceStr := c.Query("min_price"); minPriceStr != "" {
		minPrice, _ = strconv.ParseFloat(minPriceStr, 64)
	}

	var maxPrice float64
	if maxPriceStr := c.Query("max_price"); maxPriceStr != "" {
		maxPrice, _ = strconv.ParseFloat(maxPriceStr, 64)
	}

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	products, total, err := h.service.GetProducts(search, categoryID, minPrice, maxPrice, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products: " + err.Error()})
		return
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	c.JSON(http.StatusOK, gin.H{
		"data":       products,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"total_page": totalPages,
	})
}

// GetProductByID handles GET /api/v1/products/:id
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct handles PUT /api/v1/products/:id
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.CategoryID = input.CategoryID
	product.Name = input.Name
	product.Price = input.Price
	product.Stock = input.Stock
	product.Description = input.Description

	if err := h.service.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update product: " + err.Error()})
		return
	}

	// Fetch product with preloaded category for response
	resProduct, err := h.service.GetProductByID(product.ID)
	if err == nil {
		product = resProduct
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct handles DELETE /api/v1/products/:id
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.service.DeleteProduct(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
