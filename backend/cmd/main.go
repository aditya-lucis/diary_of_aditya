package main

import (
	"log"
	"os"

	"diary-of-aditya/config"
	"diary-of-aditya/models"
	"diary-of-aditya/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	config.InitDatabase()

	// Auto migrate models
	log.Println("Running database migrations...")
	config.DB.AutoMigrate(&models.User{})
	// Add more models here as we develop
	// config.DB.AutoMigrate(&models.Client{})
	// config.DB.AutoMigrate(&models.Project{})
	// etc.

	// Setup routes
	router := routes.SetupRoutes()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Diary Of Aditya API starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
