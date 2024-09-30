package server

import (
	"coffee-track/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// Middleware for logging
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Coffee Routes
	router.GET("/coffees", handlers.GetCoffees)          // Get all coffees
	router.POST("/coffees", handlers.CreateCoffee)       // Create a coffee
	router.GET("/coffees/:id", handlers.GetCoffee)       // Get a coffee by ID
	router.PUT("/coffees/:id", handlers.UpdateCoffee)    // Update a coffee by ID
	router.DELETE("/coffees/:id", handlers.DeleteCoffee) // Delete a coffee by ID

	// CoffeeDetail endpoints
	router.POST("/coffees/:id/details", handlers.CreateCoffeeDetail) // Create a coffee detail
	router.GET("/coffees/:id/details", handlers.GetCoffeeDetails)    // Get all coffee details for a coffee

	// Recipe endpoints
	router.GET("/recipes", handlers.GetRecipes)
	router.POST("/recipes", handlers.CreateRecipe)
	router.GET("/recipes/:id", handlers.GetRecipe)
	router.PUT("/recipes/:id", handlers.UpdateRecipe)
	router.DELETE("/recipes/:id", handlers.DeleteRecipe)

	// Brew endpoints
	router.GET("/brews", handlers.GetBrews)
	router.POST("/brews", handlers.CreateBrew)
	router.GET("/brews/:id", handlers.GetBrew)
	router.PUT("/brews/:id", handlers.UpdateBrew)
	router.DELETE("/brews/:id", handlers.DeleteBrew)

	return router
}
