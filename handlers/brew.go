package handlers

import (
	"coffee-track/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Brew Handlers
func CreateBrew(c *gin.Context) {
	var brew models.Brew
	if err := c.ShouldBindJSON(&brew); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&brew).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brew)
}

func GetBrews(c *gin.Context) {
	var brews []models.Brew
	if err := models.DB.Find(&brews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brews)
}

func GetBrew(c *gin.Context) {
	id := c.Param("id")
	var brew models.Brew
	if err := models.DB.First(&brew, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brew not found"})
		return
	}
	c.JSON(http.StatusOK, brew)
}

func UpdateBrew(c *gin.Context) {
	id := c.Param("id")
	var brew models.Brew
	if err := models.DB.First(&brew, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brew not found"})
		return
	}
	if err := c.ShouldBindJSON(&brew); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Save(&brew).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brew)
}

func DeleteBrew(c *gin.Context) {
	id := c.Param("id")
	if err := models.DB.Delete(&models.Brew{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
