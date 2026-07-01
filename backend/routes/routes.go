package routes

import (
	"diary-of-aditya/handlers"
	"diary-of-aditya/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Health check
	router.GET("/api/health", handlers.HealthCheck)

	// API v1
	v1 := router.Group("/api/v1")
	{
		// Public routes
		auth := v1.Group("/auth")
		{
			auth.POST("/login", handlers.Login)
			auth.POST("/register", handlers.Register)
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.JWTAuthMiddleware())
		{
			protected.GET("/profile", handlers.GetProfile)

			// Future routes will be added here
			// - Clients
			// - Projects
			// - Milestones
			// - Tasks
			// - Timesheets
			// - Invoices
			// - Expenses
		}
	}

	return router
}
