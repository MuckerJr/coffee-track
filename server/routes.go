package server

import (
	"coffee-track/handlers"
	"coffee-track/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// Middleware for logging
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Public routes
	router.GET("/login", handlers.Login)
	router.GET("/callback", handlers.Callback)

	// Coffee Routes (Public)
	router.GET("/coffees", handlers.GetCoffees)    // Get all coffees
	router.GET("/coffees/:id", handlers.GetCoffee) // Get a coffee by ID

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.JWTMiddleware())
	{
		// Coffee Routes (Protected)
		protected.POST("/coffees", handlers.CreateCoffee)       // Create a coffee
		protected.PUT("/coffees/:id", handlers.UpdateCoffee)    // Update a coffee by ID
		protected.DELETE("/coffees/:id", handlers.DeleteCoffee) // Delete a coffee by ID

		// CoffeeDetail endpoints
		protected.POST("/coffees/:id/details", handlers.CreateCoffeeDetail) // Create a coffee detail
		protected.GET("/coffees/:id/details", handlers.GetCoffeeDetails)    // Get all coffee details for a coffee

		// Recipe endpoints
		protected.GET("/recipes", handlers.GetRecipes)
		protected.POST("/recipes", handlers.CreateRecipe)
		protected.GET("/recipes/:id", handlers.GetRecipe)
		protected.PUT("/recipes/:id", handlers.UpdateRecipe)
		protected.DELETE("/recipes/:id", handlers.DeleteRecipe)

		// Brew endpoints
		protected.GET("/brews", handlers.GetBrews)
		protected.POST("/brews", handlers.CreateBrew)
		protected.GET("/brews/:id", handlers.GetBrew)
		protected.PUT("/brews/:id", handlers.UpdateBrew)
		protected.DELETE("/brews/:id", handlers.DeleteBrew)

		// Inventory endpoints
		protected.GET("/inventory/:user_id", handlers.GetInventory)
		protected.POST("/inventory", handlers.AddToInventory)
	}

	return router
}
