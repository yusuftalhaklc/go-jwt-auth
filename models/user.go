package models

// User represents a user model with username and password fields.
type User struct {
	Username string `json:"username"` // Username field
	Password string `json:"password"` // Password field
}

// users is a slice of User objects representing the available users.
var users = []User{
	{
		Username: "admin",
		Password: "1234",
	},
}

// IsThereUser checks if a user with the given username and password exists in the users slice.
func IsThereUser(user *User) bool {
	for _, v := range users {
		if v.Username == user.Username && v.Password == user.Password {
			return true
		}
	}
	return false
}
