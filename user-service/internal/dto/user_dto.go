package dto

import "time"

type RegisterUserRequest struct {
	FullName string `json:"full_name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        string
	FullName  string
	Email     string
	CreatedAt time.Time
}

type LoginRequest struct {
	Email    string
	Password string
}
