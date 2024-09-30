package handlers

import (
	"coffee-track/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CoffeeDetail Handlers
func CreateCoffeeDetail(c *gin.Context) {
	coffeeID := c.Param("id")
	var coffeeDetail models.CoffeeDetail
	if err := c.ShouldBindJSON(&coffeeDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parsedID, err := strconv.ParseUint(coffeeID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	coffeeDetail.CoffeeID = uint(parsedID)
	if err := models.DB.Create(&coffeeDetail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the quantity of the coffee in the coffee table
	var coffee models.Coffee
	if err := models.DB.First(&coffee, coffeeDetail.CoffeeID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	coffee.Quantity++
	if err := models.DB.Save(&coffee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, coffeeDetail)
}

func GetCoffeeDetails(c *gin.Context) {
	coffeeID := c.Param("id")
	var coffeeDetails []models.CoffeeDetail
	if err := models.DB.Where("coffee_id = ?", coffeeID).Find(&coffeeDetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coffeeDetails)
}
