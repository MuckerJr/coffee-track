package cli

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"coffee-track/models"
)

func listCoffees() {
	fmt.Println("Listing coffees...")
	db, err := gorm.Open(sqlite.Open("coffee.db"), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	var coffees []models.Coffee
	db.Find(&coffees)
	for _, coffee := range coffees {
		fmt.Printf("ID: %d, Name: %s, Quantity: %s\n", coffee.ID, coffee.Name, coffee.Quantity)
	}
}
