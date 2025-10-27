package auth

import (
	"time"
)

type EmailPassword struct {
	Email    string `json:"email" example:"john.doe@example.com" doc:"The email of the user"`
	Password string `json:"password" example:"securepassword123" doc:"The password of the user"`
}

type PersonName struct {
	FirstName string `json:"first_name" example:"John" doc:"The first name of the user"`
	LastName  string `json:"last_name" example:"Doe" doc:"The last name of the user"`
}

type DateAudit struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Request struct {
	EmailPassword
}

type SignupRequest struct {
	EmailPassword
	PersonName
}

type Response struct {
	Token   string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." doc:"JWT token for authenticated sessions"`
	Message string `json:"message" example:"User registered successfully" doc:"A message indicating the result of the operation"`
}
