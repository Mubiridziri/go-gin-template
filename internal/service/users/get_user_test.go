package users

import (
	"app/internal/dto"
	"app/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockGetUserRepository struct {
	mock.Mock
}

func (m *MockGetUserRepository) GetUserById(id int) (entity.User, error) {
	args := m.Called()
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockGetUserRepository) GetUserByUsername(username string) (entity.User, error) {
	args := m.Called()
	return args.Get(0).(entity.User), args.Error(1)
}

func TestGetUserById(t *testing.T) {
	mockRepo := new(MockGetUserRepository)

	defaultTime := time.Date(2025, time.January, 12, 0, 0, 0, 0, time.UTC)
	mockUser := entity.User{
		ID:        1,
		Name:      "John Doe",
		Username:  "johndoe",
		Role:      "user",
		Email:     "johndoe@example.com",
		Password:  "123456",
		CreatedAt: defaultTime,
		UpdatedAt: defaultTime,
	}

	mockRepo.On("GetUserById").Return(mockUser, nil)

	useCase := GetUserUseCase{mockRepo}

	result, err := useCase.GetUserById(1)

	assert.NoError(t, err)

	expectedResponseDto := dto.UserResponseDto{
		Name:      "John Doe",
		Username:  "johndoe",
		Role:      "user",
		Email:     "johndoe@example.com",
		CreatedAt: defaultTime,
		UpdatedAt: defaultTime,
	}

	assert.Equal(t, expectedResponseDto, result)

	mockRepo.AssertExpectations(t)
}

func TestGetUserByUsername(t *testing.T) {
	mockRepo := new(MockGetUserRepository)

	defaultTime := time.Date(2025, time.January, 12, 0, 0, 0, 0, time.UTC)
	mockUser := entity.User{
		ID:        1,
		Name:      "John Doe",
		Username:  "johndoe",
		Role:      "user",
		Email:     "johndoe@example.com",
		Password:  "123456",
		CreatedAt: defaultTime,
		UpdatedAt: defaultTime,
	}

	mockRepo.On("GetUserByUsername").Return(mockUser, nil)

	useCase := GetUserUseCase{mockRepo}

	result, err := useCase.GetUserByUsername("johndoe")

	assert.NoError(t, err)

	expectedResponseDto := dto.UserResponseDto{
		Name:      "John Doe",
		Username:  "johndoe",
		Role:      "user",
		Email:     "johndoe@example.com",
		CreatedAt: defaultTime,
		UpdatedAt: defaultTime,
	}

	assert.Equal(t, expectedResponseDto, result)

	mockRepo.AssertExpectations(t)
}
