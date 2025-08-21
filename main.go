package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"ambridge-backend/database"
	"ambridge-backend/middleware"
	"ambridge-backend/models"
	"ambridge-backend/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Initialize database connection
	database.InitDB()

	// Run auto migrations
	autoMigrate()

	// Set up Gin router
	router := gin.Default()

	// Use custom middlewares
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.LoggerMiddleware())

	// Set up routes
	routes.SetupAuthRoutes(router)
	routes.SetupProjectRoutes(router)
	routes.SetupCrewRoutes(router)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// autoMigrate runs database migrations for all models
func autoMigrate() {
	log.Println("Running database migrations...")
	err := database.DB.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.Crew{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Database migrations completed successfully")
}
