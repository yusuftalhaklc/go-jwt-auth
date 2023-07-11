package models

// Response represents a general response structure
type Response struct {
	Message string `json:"message"` // Message contains the response message
	Status  int    `json:"status"`  // Status represents the response status code
}

// LoginResponse represents the response structure for login
type LoginResponse struct {
	Token     string `json:"access_token"` // Token contains the access token for authentication
	TokenType string `json:"token_type"`   // TokenType indicates the type of the access token
}
