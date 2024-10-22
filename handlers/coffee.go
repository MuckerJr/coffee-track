package handlers

import (
	"coffee-track/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CreateCoffee creates a new coffee type blueprint
func createCoffee(c *gin.Context) {
	var newCoffee models.Coffee
	if err := c.ShouldBindJSON(&newCoffee); err != nil {
		var validationErrors gin.H
		if errs, ok := err.(validator.ValidationErrors); ok {
			validationErrors = gin.H{}
			for _, e := range errs {
				validationErrors[e.Field()] = e.Tag()
			}
		} else {
			validationErrors = gin.H{"error": err.Error()}
		}
		c.JSON(http.StatusBadRequest, validationErrors)
		return
	}

	if err := models.DB.Create(&newCoffee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newCoffee)
}

// AddToInventory adds a quantity of coffee to user's inventory
func addToInventory(c *gin.Context) {
	var newItem models.InventoryItem
	if err := c.ShouldBindJSON(&newItem); err != nil {
		var validationErrors gin.H
		if errs, ok := err.(validator.ValidationErrors); ok {
			validationErrors = make(gin.H)
			for _, e := range errs {
				validationErrors[e.Field()] = e.Error()
			}
		} else {
			validationErrors = gin.H{"error": err.Error()}
		}
		c.JSON(http.StatusBadRequest, validationErrors)
		return
	}

	var existingItem models.InventoryItem
	if err := models.DB.Where("user_id =? AND coffee_id = ?", newItem.UserId, newItem.CoffeeID).First(&existingItem).Error; err != nil {
		// if the item exists, update quant, sizes and grinds
		existingItem.Quantity += newItem.Quantity
		existingItem.Sizes = mergeStrings(existingItem.Sizes, newItem.Sizes)
		existingItem.Grinds = mergeGrinds(existingItem.Grinds, newItem.Grinds)

		if err := models.DB.Save(&existingItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, existingItem)
		return
	}

	//Item does not exist
	if err := models.DB.Create(&newItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newItem)
}

// GetInventory returns a user's inv
func GetInventory(c *gin.Context) {
	userID := c.Param("user_id")
	var inventory []models.InventoryItem
	if err := models.DB.Preload("Coffee").Where("user_id = ?", userID).Find(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, inventory)
}

// mergeString mergers two slices of string, removing duplicates
func mergeStrings(existing, new []string) []string {
	existingMap := make(map[string]bool)
	for _, s := range existing {
		existingMap[s] = true
	}
	for _, s := range new {
		if !existingMap[s] {
			existing = append(existing, s)
		}
	}
	return existing
}

// mergeGrinds merges two slices of Grind, removing duplicates
func mergeGrinds(existing, new []models.Grind) []models.Grind {
	existingMap := make(map[models.Grind]bool)
	for _, g := range existing {
		existingMap[g] = true
	}
	for _, g := range new {
		if !existingMap[g] {
			existing = append(existing, g)
		}
	}
	return existing
}
