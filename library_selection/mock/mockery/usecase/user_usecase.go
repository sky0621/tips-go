package usecase

import (
	"context"

	"github.com/sky0621/tips-go/library_selection/mock/mockery/domain"
)

// UserUseCase defines the interface for user business logic
//go:generate mockery --name UserUseCase --output ../mocks/usecase
type UserUseCase interface {
	// GetUser retrieves a user by ID
	GetUser(ctx context.Context, id string) (*domain.User, error)

	// GetAllUsers retrieves all users
	GetAllUsers(ctx context.Context) ([]*domain.User, error)

	// CreateUser creates a new user
	CreateUser(ctx context.Context, name, email string) (*domain.User, error)

	// UpdateUser updates an existing user
	UpdateUser(ctx context.Context, id, name, email string) (*domain.User, error)

	// DeleteUser removes a user by ID
	DeleteUser(ctx context.Context, id string) error
}

// userUseCase implements UserUseCase
type userUseCase struct {
	userRepo domain.UserRepository
}

// NewUserUseCase creates a new UserUseCase
func NewUserUseCase(userRepo domain.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

// GetUser retrieves a user by ID
func (uc *userUseCase) GetUser(ctx context.Context, id string) (*domain.User, error) {
	return uc.userRepo.FindByID(ctx, id)
}

// GetAllUsers retrieves all users
func (uc *userUseCase) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	return uc.userRepo.FindAll(ctx)
}

// CreateUser creates a new user
func (uc *userUseCase) CreateUser(ctx context.Context, name, email string) (*domain.User, error) {
	user, err := domain.NewUser(name, email)
	if err != nil {
		return nil, err
	}

	err = uc.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser updates an existing user
func (uc *userUseCase) UpdateUser(ctx context.Context, id, name, email string) (*domain.User, error) {
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user.Name = name
	user.Email = email

	if err := user.Validate(); err != nil {
		return nil, err
	}

	err = uc.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser removes a user by ID
func (uc *userUseCase) DeleteUser(ctx context.Context, id string) error {
	return uc.userRepo.Delete(ctx, id)
}
