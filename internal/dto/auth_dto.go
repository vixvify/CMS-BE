package dto

import "github.com/google/uuid"

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}