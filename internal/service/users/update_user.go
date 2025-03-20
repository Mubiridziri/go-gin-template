package users

import (
	"app/internal/dto"
	"app/internal/entity"
)

type UpdateUserRepository interface {
	GetUserById(id int) (entity.User, error)
	UpdateUser(user *entity.User) error
}

type UpdateUserUseCase struct {
	UpdateUserRepository
}

func (useCase UpdateUserUseCase) UpdateUser(id int, input dto.SaveUserDto) (dto.UserResponseDto, error) {

	user, err := useCase.UpdateUserRepository.GetUserById(id)

	if err != nil {
		return dto.UserResponseDto{}, err
	}

	user.Name = input.Name
	user.Username = input.Username
	user.Email = input.Email

	if input.Password != "" {
		user.Password = input.Password
	}

	if err := useCase.UpdateUserRepository.UpdateUser(&user); err != nil {
		return dto.UserResponseDto{}, err
	}

	return fromDBUser(&user), nil
}
