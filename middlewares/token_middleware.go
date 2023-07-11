package middleware

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yusuftalhaklc/go-jwt-auth/models"
	"github.com/yusuftalhaklc/go-jwt-auth/utils"
)

// TokenMiddleware is a middleware function that validates the JWT token in the Authorization header
func TokenMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the Authorization header from the request
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			unauthorizedStatus := http.StatusUnauthorized
			return c.Status(unauthorizedStatus).JSON(models.Response{Message: "unauthorized", Status: unauthorizedStatus})
		}
		// Extract the token from the Authorization header
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		// Verify the token and retrieve the claims
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			unauthorizedStatus := http.StatusUnauthorized
			return c.Status(unauthorizedStatus).JSON(models.Response{Message: err.Error(), Status: unauthorizedStatus})
		}
		// Store the claims in the context locals for future use
		c.Locals(claims)
		// Proceed to the next middleware or route handler
		return c.Next()
	}
}
