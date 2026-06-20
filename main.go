package main

import (
	"fmt"
	"go-sqlite-crud/config"
	"go-sqlite-crud/routes"
)

func main() {
	// Initialize database and migration
	config.InitDB()

	// Setup routing structure by injecting the database connection
	r := routes.SetupRouter(config.DB)

	fmt.Println("Server is running on port :8080")
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
