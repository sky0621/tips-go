package domain

import "context"

// UserRepository defines the interface for user data access
//go:generate mockery --name UserRepository --output ../mocks/domain
type UserRepository interface {
	// FindByID retrieves a user by ID
	FindByID(ctx context.Context, id string) (*User, error)
	
	// FindAll retrieves all users
	FindAll(ctx context.Context) ([]*User, error)
	
	// Create saves a new user
	Create(ctx context.Context, user *User) error
	
	// Update updates an existing user
	Update(ctx context.Context, user *User) error
	
	// Delete removes a user by ID
	Delete(ctx context.Context, id string) error
}