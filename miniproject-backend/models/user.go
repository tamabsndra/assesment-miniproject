package models

import "time"

type User struct {
    ID        uint      `json:"id"`
    Email     string    `json:"email" validate:"required,email"`
    Password  string    `json:"password,omitempty" validate:"required,min=6"`
    Name      string    `json:"name" validate:"required"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type VerifyUser struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email" validate:"required,email"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserInPost struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email" validate:"required,email"`
	Name      string    `json:"name" validate:"required"`
}

type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
    Token string `json:"token"`
    User  User   `json:"user"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
}

type RegisterResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
