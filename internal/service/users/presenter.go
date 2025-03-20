package users

import (
	"app/internal/dto"
	"app/internal/entity"
)

func fromDBUser(user *entity.User) dto.UserResponseDto {
	return dto.UserResponseDto{
		Name:      user.Name,
		Username:  user.Username,
		Role:      user.Role,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
