package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	r "github.com/yusuftalhaklc/go-jwt-auth/router"
)

func main() {
	app := fiber.New()
	// Use logger middleware for logging requests
	app.Use(logger.New())
	// Use cors middleware for handling CORS
	app.Use(cors.New())
	// Set up routes
	r.SetupRoutes(app)
	// Start the application and listen on port 8080
	app.Listen(":8080")
}
