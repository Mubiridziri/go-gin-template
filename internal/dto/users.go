package dto

import "time"

type UserResponseDto struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SaveUserDto struct {
	Name     string `json:"name"  binding:"required"`
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Email    string `json:"email"  binding:"required"`
}

type PaginatedUsersList struct {
	Total   int64             `json:"total"`
	Entries []UserResponseDto `json:"entries"`
}
