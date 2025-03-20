package users

import (
	"app/internal/dto"
	"app/internal/entity"
	"errors"
)

type SecurityRepository interface {
	GetUserByUsername(username string) (entity.User, error)
}

type SecurityUseCase struct {
	SecurityRepository
}

func (us SecurityUseCase) LoginUser(input dto.UserLogin) (dto.UserResponseDto, error) {
	user, err := us.SecurityRepository.GetUserByUsername(input.Username)

	if err != nil {
		return dto.UserResponseDto{}, err
	}

	if !user.IsPasswordCorrect(input.Password) {
		return dto.UserResponseDto{}, errors.New("invalid credentials")
	}

	return fromDBUser(&user), nil
}
