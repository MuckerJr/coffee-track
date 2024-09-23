package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Coffee struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Recipe struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

func main() {
	r := gin.Default()
	db, err := gorm.Open(sqlite.Open("coffee.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Coffee{}, &Recipe{})
	r.GET("/coffees", func(c *gin.Context) {
		var coffees []Coffee
		db.Find(&coffees)
		c.JSON(200, coffees)
	})

	r.POST("/coffees", func(c *gin.Context) {
		var coffee Coffee
		if err := c.ShouldBindJSON(&coffee); err == nil {
			db.Create(&coffee)
			c.JSON(200, coffee)
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
		}
	})

	r.GET("/recipes", func(c *gin.Context) {
		var recipes []Recipe
		db.Find(&recipes)
		c.JSON(200, recipes)
	})

	r.POST("/recipes", func(c *gin.Context) {
		var recipe Recipe
		if err := c.ShouldBindJSON(&recipe); err == nil {
			db.Create(&recipe)
			c.JSON(200, recipe)
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
		}
	})

	r.Run()
}
