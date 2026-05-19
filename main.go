package main

import (
	"log"
	"os"

	"backend/config"
	"backend/url"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialize environment variables
	config.InitConfig()

	// Initialize MongoDB connection
	config.InitMongoDB()

	// Initialize Fiber app
	app := fiber.New()

	// Use CORS middleware
	app.Use(cors.New(config.CorsConfig()))

	// Static files for frontend
	app.Static("/", "./frontend")

	// Register routes
	url.Web(app)

	// Determine port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	err := app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
