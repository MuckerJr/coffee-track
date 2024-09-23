//go:build !cli
// +build !cli

package main

import (
	"coffee-track/models"

	"github.com/gin-gonic/gin"
)

func main() {
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

	r.GET("/recipes", func(c *gin.Context) {
		var recipes []models.Recipe
		models.DB.Find(&recipes)
		c.JSON(200, recipes)
	})

	r.POST("/recipes", func(c *gin.Context) {
		var recipe models.Recipe
		if err := c.ShouldBindJSON(&recipe); err == nil {
			models.DB.Create(&recipe)
			c.JSON(200, recipe)
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
		}
	})

	r.Run()
}
