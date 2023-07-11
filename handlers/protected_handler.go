package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusuftalhaklc/go-jwt-auth/models"
)

func Protected(c *fiber.Ctx) error {
	// Create a response with the success message and status code
	response := models.Response{Message: successMessage, Status: status200}
	// Return the response as JSON with the corresponding status code
	return c.Status(status200).JSON(response)
}
