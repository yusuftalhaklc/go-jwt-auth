package handlers

import "net/http"

// unauthorized represents the unauthorized message
var unauthorized = "unauthorized"

// status401 represents the HTTP status code for unauthorized
var status401 = http.StatusUnauthorized

// successMessage represents the success message
var successMessage = "Welcome to the protected area"

// status200 represents the HTTP status code for success
var status200 = http.StatusOK

// badRequestMessage represents the message for a bad request
var badRequestMessage = "Something's wrong with your input"

// status400 represents the HTTP status code for a bad request
var status400 = http.StatusBadRequest

// status500 represents the HTTP status code for internal server error
var status500 = http.StatusInternalServerError
