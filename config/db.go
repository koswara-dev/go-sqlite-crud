package config

import (
	"fmt"
	"go-sqlite-crud/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("crud.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	fmt.Println("Database connection established.")

	// Auto Migration
	err = DB.AutoMigrate(&models.Category{}, &models.Product{})
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
	fmt.Println("Database migration completed.")
}
