package users

import (
	"app/internal/dto"
	"app/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockGetUsersRepository struct {
	mock.Mock
}

func (m *MockGetUsersRepository) ListUsers(page, limit int) ([]entity.User, error) {
	args := m.Called(page, limit)
	return args.Get(0).([]entity.User), args.Error(1)
}

func (m *MockGetUsersRepository) GetUsersCount() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

func TestListUsers(t *testing.T) {
	mockRepo := new(MockGetUsersRepository)

	mockUsers := []entity.User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	mockRepo.On("ListUsers", 1, 10).Return(mockUsers, nil)
	mockRepo.On("GetUsersCount").Return(int64(2))

	useCase := GetUsersListUseCase{mockRepo}

	result, err := useCase.ListUsers(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), result.Total)

	expectedEntries := []dto.UserResponseDto{
		fromDBUser(&mockUsers[0]),
		fromDBUser(&mockUsers[1]),
	}

	assert.Equal(t, expectedEntries, result.Entries)

	mockRepo.AssertExpectations(t)

}
