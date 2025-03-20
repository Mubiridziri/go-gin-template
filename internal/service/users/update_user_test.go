package users

import (
	"app/internal/dto"
	"app/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockUpdateUserRepository struct {
	mock.Mock
}

func (m *MockUpdateUserRepository) GetUserById(id int) (entity.User, error) {
	args := m.Called()
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUpdateUserRepository) UpdateUser(user *entity.User) error {
	args := m.Called()
	return args.Error(0)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockUpdateUserRepository)

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
	mockRepo.On("UpdateUser").Return(nil)

	useCase := UpdateUserUseCase{mockRepo}

	saveDto := dto.SaveUserDto{
		Name:     "John Doe",
		Username: "johndoe",
		Password: "123456",
		Email:    "johndoe@example.com",
	}

	result, err := useCase.UpdateUser(1, saveDto)

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
