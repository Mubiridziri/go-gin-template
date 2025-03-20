package users

import (
	"app/internal/dto"
	"app/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockRemoveUserRepository struct {
	mock.Mock
}

func (m *MockRemoveUserRepository) GetUserById(id int) (entity.User, error) {
	args := m.Called()
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockRemoveUserRepository) RemoveUser(user *entity.User) error {
	args := m.Called()
	return args.Error(0)
}

func TestRemoveUser(t *testing.T) {
	mockRepo := new(MockRemoveUserRepository)

	defaultTime := time.Date(2025, time.January, 12, 0, 0, 0, 0, time.UTC)
	mockUser := entity.User{
		ID:        1,
		Name:      "John Doe",
		Username:  "johndoe",
		Role:      "admin",
		Email:     "johndoe@example.com",
		Password:  "123456",
		CreatedAt: defaultTime,
		UpdatedAt: defaultTime,
	}

	mockRepo.On("GetUserById").Return(mockUser, nil)
	mockRepo.On("RemoveUser").Return(nil)

	useCase := RemoveUserUseCase{mockRepo}

	result, err := useCase.RemoveUser(1)

	assert.NoError(t, err)

	expectedResponseDto := dto.UserResponseDto{
		Name:      "John Doe",
		Username:  "johndoe",
		Role:      "admin",
		Email:     "johndoe@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	assert.Equal(t, expectedResponseDto.Name, result.Name)
	assert.Equal(t, expectedResponseDto.Username, result.Username)
	assert.Equal(t, expectedResponseDto.Email, result.Email)
	assert.Equal(t, expectedResponseDto.Role, result.Role)

	mockRepo.AssertExpectations(t)
}
