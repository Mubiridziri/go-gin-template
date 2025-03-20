package users

import (
	"app/internal/dto"
	"app/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockCreateUserRepository struct {
	mock.Mock
}

func (m *MockCreateUserRepository) CreateUser(user *entity.User) error {
	args := m.Called()
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockCreateUserRepository)

	mockRepo.On("CreateUser").Return(nil)

	useCase := CreateUserUseCase{mockRepo}

	saveDto := dto.SaveUserDto{
		Name:     "John Doe",
		Username: "johndoe",
		Password: "123456",
		Email:    "johndoe@example.com",
	}

	result, err := useCase.CreateUser(saveDto)

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
