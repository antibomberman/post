package models

type User struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	CreatedAt string `db:"created_at"`
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
