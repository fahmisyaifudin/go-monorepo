package model

// User represents a user in the system
type User struct {
	ID               string   `json:"id"`                 // Unique identifier for the user
	Email            string  `json:"email"`              // User's email address
	FirstName        string  `json:"first_name"`         // User's first name
	LastName         *string `json:"last_name"`          // User's last name
	Password 		 string	 `json:"password"`			 // User's password 
	CreatedAt        int64   `json:"created_at"`         // Timestamp when the user was created
	UpdatedAt        *int64  `json:"updated_at"`         // Timestamp when the user was last updated
}

type LoginRequest struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}