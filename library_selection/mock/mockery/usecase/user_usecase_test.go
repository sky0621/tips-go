package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/sky0621/tips-go/library_selection/mock/mockery/domain"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	NewMockUserUseCase(t)
	mockRepo := new(domain.NewMockUserUseCase)
	uc := NewUserUseCase(mockRepo)
	ctx := context.Background()
	user := &domain.User{ID: "1", Name: "test", Email: "test@example.com"}

	mockRepo.On("FindByID", ctx, "1").Return(user, nil)

	got, err := uc.GetUser(ctx, "1")
	assert.NoError(t, err)
	assert.Equal(t, user, got)
	mockRepo.AssertExpectations(t)
}

func TestGetAllUsers(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := usecase.NewUserUseCase(mockRepo)
	ctx := context.Background()
	users := []*domain.User{
		{ID: "1", Name: "test1", Email: "test1@example.com"},
		{ID: "2", Name: "test2", Email: "test2@example.com"},
	}

	mockRepo.On("FindAll", ctx).Return(users, nil)

	got, err := uc.GetAllUsers(ctx)
	assert.NoError(t, err)
	assert.Equal(t, users, got)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := usecase.NewUserUseCase(mockRepo)
	ctx := context.Background()
	name := "test"
	email := "test@example.com"
	user, _ := domain.NewUser(name, email)

	mockRepo.On("Create", ctx, user).Return(nil)

	got, err := uc.CreateUser(ctx, name, email)
	assert.NoError(t, err)
	assert.Equal(t, user, got)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := usecase.NewUserUseCase(mockRepo)
	ctx := context.Background()
	id := "1"
	name := "updated"
	email := "updated@example.com"
	user := &domain.User{ID: id, Name: "old", Email: "old@example.com"}

	mockRepo.On("FindByID", ctx, id).Return(user, nil)
	mockRepo.On("Update", ctx, user).Return(nil)

	got, err := uc.UpdateUser(ctx, id, name, email)
	assert.NoError(t, err)
	assert.Equal(t, name, got.Name)
	assert.Equal(t, email, got.Email)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := usecase.NewUserUseCase(mockRepo)
	ctx := context.Background()
	id := "1"

	mockRepo.On("Delete", ctx, id).Return(nil)

	err := uc.DeleteUser(ctx, id)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// 異常系の例
func TestGetUser_NotFound(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := usecase.NewUserUseCase(mockRepo)
	ctx := context.Background()

	mockRepo.On("FindByID", ctx, "notfound").Return(nil, errors.New("not found"))

	got, err := uc.GetUser(ctx, "notfound")
	assert.Error(t, err)
	assert.Nil(t, got)
	mockRepo.AssertExpectations(t)
}
