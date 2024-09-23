//go:build cli
// +build cli

package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"coffee-track/models"
)

func runApp() {
	models.InitDB()

	app := &cli.App{
		Name:  "coffee-track",
		Usage: "Track your coffee consumption and recipes",
		Commands: []*cli.Command{
			{
				Name:  "add-coffee",
				Usage: "Add a new coffee to the inventory",
				Action: func(c *cli.Context) error {
					db, err := gorm.Open(sqlite.Open("coffee.db"), &gorm.Config{})
					if err != nil {
						return err
						// Should I return a panic?
					}
					db.AutoMigrate(&models.Coffee{})
					coffee := models.Coffee{Name: c.Args().Get(0), Quantity: c.Int("quantity")}
					db.Create(&coffee)
					fmt.Printf("Coffee added: %s\n", coffee.Name)
					return nil
				},
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "quantity",
						Value: 1,
						Usage: "Quantity of coffee to add",
					},
				},
			},
			{
				Name:  "list-coffees",
				Usage: "List all coffees in the inventory",
				Action: func(c *cli.Context) error {
					db, err := gorm.Open(sqlite.Open("coffee.db"), &gorm.Config{})
					if err != nil {
						return err
					}
					var coffees []models.Coffee
					db.Find(&coffees)
					for _, coffee := range coffees {
						fmt.Printf("ID: %d, Name: %s, Quantity: %d\n", coffee.ID, coffee.Name, coffee.Quantity)
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	runApp()
}
