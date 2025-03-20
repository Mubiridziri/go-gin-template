package users

import (
	"app/internal/dto"
	"app/internal/entity"
)

type RemoveUserRepository interface {
	GetUserById(id int) (entity.User, error)
	RemoveUser(user *entity.User) error
}

type RemoveUserUseCase struct {
	RemoveUserRepository
}

func (useCase RemoveUserUseCase) RemoveUser(id int) (dto.UserResponseDto, error) {

	user, err := useCase.RemoveUserRepository.GetUserById(id)

	if err != nil {
		return dto.UserResponseDto{}, err
	}

	if err := useCase.RemoveUserRepository.RemoveUser(&user); err != nil {
		return dto.UserResponseDto{}, err
	}

	return fromDBUser(&user), nil
}
