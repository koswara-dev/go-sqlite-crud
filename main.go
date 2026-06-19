package main

import (
	"fmt"
	"go-sqlite-crud/config"
	"go-sqlite-crud/routes"
)

func main() {
	// Initialize database and migration
	config.InitDB()

	// Setup routing structure
	r := routes.SetupRouter()

	fmt.Println("Server is running on port :8080")
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
