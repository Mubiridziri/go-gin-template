package users

import (
	"app/internal/dto"
	"app/internal/entity"
	"time"
)

type CreateUserRepository interface {
	CreateUser(user *entity.User) error
}

type CreateUserUseCase struct {
	CreateUserRepository
}

func (useCase CreateUserUseCase) CreateUser(input dto.SaveUserDto) (dto.UserResponseDto, error) {
	user := entity.User{
		Name:      input.Name,
		Username:  input.Username,
		Password:  input.Password,
		Email:     input.Email,
		Role:      "admin", //TODO temp
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := useCase.CreateUserRepository.CreateUser(&user); err != nil {
		return dto.UserResponseDto{}, err
	}

	return fromDBUser(&user), nil
}
