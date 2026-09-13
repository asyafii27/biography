package main

import (
	"biography-api/config"
	"biography-api/controllers"
	"biography-api/migrations"
	"biography-api/repositories"
	"biography-api/routes"
	"biography-api/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to Database
	config.ConnectDB()

	// Run Migrations
	migrations.Migrate()

	// Setup Dependencies
	db := config.DB
	experienceRepo := repositories.NewExperienceRepository(db)
	experienceService := services.NewExperienceService(experienceRepo)
	experienceController := controllers.NewExperienceController(experienceService)

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	// Setup Gin Router
	router := gin.Default()

	// Setup all routes (prefix /api is handled inside this function)
	routes.SetupRoutes(router, experienceController, authController)

	// Get Port from .env
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port: %s\n", port)
	
	// Start Server
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
