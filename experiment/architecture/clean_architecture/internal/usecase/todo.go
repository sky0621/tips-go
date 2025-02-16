package usecase

import (
	"context"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture/internal/domain"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *domain.Todo) error
	GetTodos(ctx context.Context) ([]*domain.Todo, error)
	GetTodoByID(ctx context.Context, id int64) (*domain.Todo, error)
	UpdateTodo(ctx context.Context, todo *domain.Todo) error
	DeleteTodo(ctx context.Context, id int64) error
}

type TodoUsecase interface {
	CreateTodo(ctx context.Context, title string) (*domain.Todo, error)
	ListTodos(ctx context.Context) ([]*domain.Todo, error)
	GetTodo(ctx context.Context, id int64) (*domain.Todo, error)
	UpdateTodo(ctx context.Context, id int64, title string, completed bool) (*domain.Todo, error)
	DeleteTodo(ctx context.Context, id int64) error
}

type todoUsecase struct {
	repo TodoRepository
}

func NewTodoUsecase(repo TodoRepository) TodoUsecase {
	return &todoUsecase{repo: repo}
}

func (u *todoUsecase) CreateTodo(ctx context.Context, title string) (*domain.Todo, error) {
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

func (u *todoUsecase) ListTodos(ctx context.Context) ([]*domain.Todo, error) {
	return u.repo.GetTodos(ctx)
}

func (u *todoUsecase) GetTodo(ctx context.Context, id int64) (*domain.Todo, error) {
	return u.repo.GetTodoByID(ctx, id)
}

func (u *todoUsecase) UpdateTodo(ctx context.Context, id int64, title string, completed bool) (*domain.Todo, error) {
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

func (u *todoUsecase) DeleteTodo(ctx context.Context, id int64) error {
	return u.repo.DeleteTodo(ctx, id)
}
