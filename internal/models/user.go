package models

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt string
}

type UserCreate struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
