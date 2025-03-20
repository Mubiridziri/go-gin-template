package users

import (
	"app/internal/dto"
	"app/internal/entity"
)

type GetUserRepository interface {
	GetUserById(id int) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
}

type GetUserUseCase struct {
	GetUserRepository
}

func (useCase GetUserUseCase) GetUserById(id int) (dto.UserResponseDto, error) {

	user, err := useCase.GetUserRepository.GetUserById(id)

	if err != nil {
		return dto.UserResponseDto{}, err
	}

	return fromDBUser(&user), nil
}

func (useCase GetUserUseCase) GetUserByUsername(username string) (dto.UserResponseDto, error) {
	user, err := useCase.GetUserRepository.GetUserByUsername(username)

	if err != nil {
		return dto.UserResponseDto{}, err
	}

	return fromDBUser(&user), nil
}
