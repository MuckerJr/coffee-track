package handlers

import (
	"coffee-track/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Coffee Handlers
func GetCoffees(c *gin.Context) {
	var coffees []models.Coffee
	if err := models.DB.Preload("Details").Find(&coffees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coffees)
}

func CreateCoffee(c *gin.Context) {
	var newCoffee models.Coffee
	if err := c.ShouldBindJSON(&newCoffee); err != nil {
		// If binding fails, respond with a 400 and detailed error messages
		var validationErrors gin.H
		if errs, ok := err.(validator.ValidationErrors); ok {
			validationErrors = make(gin.H)
			for _, fieldErr := range errs {
				validationErrors[fieldErr.Field()] = fieldErr.Error()
			}
		} else {
			validationErrors = gin.H{"json": err.Error()}
		}
		c.JSON(http.StatusBadRequest, validationErrors)
		return
	}

	// No errors, proceed to create the coffee
	var existingCoffee models.Coffee
	if err := models.DB.Preload("Details").Where("name = ? AND vendor = ?", newCoffee.Name, newCoffee.Vendor).First(&existingCoffee).Error; err == nil {
		// If the coffee already exists, increment the quantity and merge details
		existingCoffee.Quantity += newCoffee.Quantity

		for _, newDetail := range newCoffee.Details {
			found := false
			for i, existingDetail := range existingCoffee.Details {
				if existingDetail.Grind == newDetail.Grind && existingDetail.Roast == newDetail.Roast {
					// Merge sizes if grind and roast match (maybe not)
					existingCoffee.Details[i].Size = mergeSizes(existingDetail.Size, newDetail.Size)
					found = true
					break
				}
			}
			if !found {
				// Add new detail if not found
				existingCoffee.Details = append(existingCoffee.Details, newDetail)
			}
		}

		if err := models.DB.Save(&existingCoffee).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, existingCoffee)
		return
	} else {
		// Coffee does not exist, create a new entry
		if err := models.DB.Create(&newCoffee).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, newCoffee)
	}

}

func mergeSizes(existingSizes, newSize []string) []string {
	for i, size := range existingSizes {
		if size == newSize[i] {
			return existingSizes
		}
	}
	return append(existingSizes, newSize[0])
}

func GetCoffee(c *gin.Context) {
	id := c.Param("id")
	var coffee models.Coffee
	if err := models.DB.Preload("Details").First(&coffee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coffee not found"})
		return
	}
	c.JSON(http.StatusOK, coffee)
}

func UpdateCoffee(c *gin.Context) {
	id := c.Param("id")
	var coffee models.Coffee
	if err := models.DB.First(&coffee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coffee not found"})
		return
	}
	if err := c.ShouldBindJSON(&coffee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Save(&coffee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coffee)
}

func DeleteCoffee(c *gin.Context) {
	id := c.Param("id")
	if err := models.DB.Delete(&models.Coffee{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
