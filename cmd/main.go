package main

import (
	"coffee-track/models"
	"coffee-track/server"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize the database
	var err error
	models.DB, err = gorm.Open(sqlite.Open("coffee.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Migrate the schema
	err = models.DB.AutoMigrate(&models.Coffee{}, &models.CoffeeDetail{}, &models.Recipe{}, &models.Brew{})
	if err != nil {
		log.Fatalf("Failed to migrate schema: %v", err)
	}

	// Initialize the database
	models.InitDB()
	router := server.InitRouter()

	// Log a message indicating the server is running
	log.Println("Server running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
