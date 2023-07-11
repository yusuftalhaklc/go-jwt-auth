package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusuftalhaklc/go-jwt-auth/handlers"
	middlewares "github.com/yusuftalhaklc/go-jwt-auth/middlewares"
)

func SetupRoutes(app *fiber.App) {
	// Create a new group under "/api" route
	api := app.Group("/api")
	// Create a new subgroup under "/api/v1" route
	v1 := api.Group("/v1")
	// Create a token middleware for authentication
	tokenMiddleware := middlewares.TokenMiddleware()
	// Handle POST requests to "/api/v1/Login" route with the Login handler
	v1.Post("/Login", handlers.Login)
	// Handle GET requests to "/api/v1/protected" route with the Protected handler, and apply the token middleware for authentication
	v1.Get("/protected", tokenMiddleware, handlers.Protected)
}
