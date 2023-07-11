package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/yusuftalhaklc/go-jwt-auth/config"
	"github.com/yusuftalhaklc/go-jwt-auth/models"
)

// secretKey is the secret key used for signing the JWT token.
var secretKey = []byte(config.Config("SECRET"))

// CreateToken creates a new JWT token for the given user.
func CreateToken(user *models.User) (string, error) {
	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the token claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies the authenticity of the given JWT token and returns its claims.
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	// Check for parsing errors or invalid tokens
	if err != nil || !token.Valid {
		return nil, errors.New("unauthorized")
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid token claims")
	}

	return claims, nil
}
