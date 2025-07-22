package memory

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/sky0621/tips-go/library_selection/mock/mockery/domain"
)

// UserRepository is an in-memory implementation of domain.UserRepository
type UserRepository struct {
	users map[string]*domain.User
	mutex sync.RWMutex
	nextID int
}

// NewUserRepository creates a new in-memory user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*domain.User),
		nextID: 1,
	}
}

// FindByID retrieves a user by ID
func (r *UserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// FindAll retrieves all users
func (r *UserRepository) FindAll(ctx context.Context) ([]*domain.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]*domain.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

// Create saves a new user
func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Generate ID
	user.ID = generateID(r.nextID)
	r.nextID++

	// Set timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Save user
	r.users[user.ID] = user

	return nil
}

// Update updates an existing user
func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, exists := r.users[user.ID]
	if !exists {
		return errors.New("user not found")
	}

	// Update timestamp
	user.UpdatedAt = time.Now()

	// Save user
	r.users[user.ID] = user

	return nil
}

// Delete removes a user by ID
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, exists := r.users[id]
	if !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)

	return nil
}

// Helper function to generate ID
func generateID(id int) string {
	return string(rune(id + 64)) // A=1, B=2, etc.
}