package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusuftalhaklc/go-jwt-auth/models"
	"github.com/yusuftalhaklc/go-jwt-auth/utils"
)

func Login(c *fiber.Ctx) error {
	// Create a new User model instance
	user := new(models.User)

	// Parse the request body and bind it to the User model
	err := c.BodyParser(&user)
	if err != nil {
		response := models.Response{Message: badRequestMessage, Status: status400}
		return c.Status(status400).JSON(response)
	}

	// Check if the user exists
	if !(models.IsThereUser(user)) {
		response := models.Response{Message: unauthorized, Status: status401}
		return c.Status(status401).JSON(response)
	}

	// Create a JWT token for the authenticated user
	token, err := utils.CreateToken(user)
	if err != nil {
		response := models.Response{Message: err.Error(), Status: status500}
		return c.Status(status500).JSON(response)
	}

	// Create a LoginResponse containing the token and token type
	loginResponse := models.LoginResponse{Token: token, TokenType: "bearer"}

	return c.Status(status200).JSON(loginResponse)
}
