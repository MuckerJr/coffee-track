package server

import (
	"fmt"
	"os"

	"coffee-track/models"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	fmt.Println("Starting server...")
	r := gin.Default()
	models.InitDB()

	r.GET("/coffees", func(c *gin.Context) {
		var coffees []models.Coffee
		models.DB.Find(&coffees)
		c.JSON(200, coffees)
	})

	r.POST("/coffees", func(c *gin.Context) {
		var coffee models.Coffee
		if err := c.ShouldBindJSON(&coffee); err == nil {
			models.DB.Create(&coffee)
			c.JSON(200, coffee)
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
		}
	})

	if err := r.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Server Error: %v\n", err)
		os.Exit(1)
	}
}
