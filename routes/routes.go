package routes

import (
	"go-sqlite-crud/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Root response
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Gin SQLite CRUD & Search API!",
		})
	})

	api := r.Group("/api/v1")
	{
		// Category Routes
		api.POST("/categories", handlers.CreateCategory)
		api.GET("/categories", handlers.GetCategories)
		api.GET("/categories/:id", handlers.GetCategoryByID)
		api.PUT("/categories/:id", handlers.UpdateCategory)
		api.DELETE("/categories/:id", handlers.DeleteCategory)

		// Product Routes
		api.POST("/products", handlers.CreateProduct)
		api.GET("/products", handlers.GetProducts)
		api.GET("/products/:id", handlers.GetProductByID)
		api.PUT("/products/:id", handlers.UpdateProduct)
		api.DELETE("/products/:id", handlers.DeleteProduct)
	}

	return r
}
