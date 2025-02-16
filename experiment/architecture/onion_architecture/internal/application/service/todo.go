package service

import (
	"context"
	"github.com/sky0621/tips-go/experiment/architecture/onion_architecture/internal/domain"
)

type TodoService interface {
	CreateTodo(ctx context.Context, title string) (*domain.Todo, error)
	ListTodos(ctx context.Context) ([]*domain.Todo, error)
	GetTodo(ctx context.Context, id int64) (*domain.Todo, error)
	UpdateTodo(ctx context.Context, id int64, title string, completed bool) (*domain.Todo, error)
	DeleteTodo(ctx context.Context, id int64) error
}

type todoService struct {
	repo domain.TodoRepository
}

func NewTodo(repo domain.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (u *todoService) CreateTodo(ctx context.Context, title string) (*domain.Todo, error) {
	todo := &domain.Todo{
		Title:     title,
		Completed: false,
	}
	err := u.repo.CreateTodo(ctx, todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (u *todoService) ListTodos(ctx context.Context) ([]*domain.Todo, error) {
	return u.repo.GetTodos(ctx)
}

func (u *todoService) GetTodo(ctx context.Context, id int64) (*domain.Todo, error) {
	return u.repo.GetTodoByID(ctx, id)
}

func (u *todoService) UpdateTodo(ctx context.Context, id int64, title string, completed bool) (*domain.Todo, error) {
	todo, err := u.repo.GetTodoByID(ctx, id)
	if err != nil {
		return nil, err
	}
	todo.Title = title
	todo.Completed = completed
	err = u.repo.UpdateTodo(ctx, todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (u *todoService) DeleteTodo(ctx context.Context, id int64) error {
	return u.repo.DeleteTodo(ctx, id)
}
