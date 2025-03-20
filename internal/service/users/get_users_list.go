package users

import (
	"app/internal/dto"
	"app/internal/entity"
)

type GetUsersRepository interface {
	ListUsers(page, limit int) ([]entity.User, error)
	GetUsersCount() int64
}

type GetUsersListUseCase struct {
	GetUsersRepository
}

func (useCase GetUsersListUseCase) ListUsers(page, limit int) (dto.PaginatedUsersList, error) {

	users, err := useCase.GetUsersRepository.ListUsers(page, limit)

	if err != nil {
		return dto.PaginatedUsersList{}, err
	}

	var userResponseDtos []dto.UserResponseDto

	for _, user := range users {
		userResponseDtos = append(userResponseDtos, fromDBUser(&user))
	}

	return dto.PaginatedUsersList{
		Total:   useCase.GetUsersRepository.GetUsersCount(),
		Entries: userResponseDtos,
	}, nil
}
