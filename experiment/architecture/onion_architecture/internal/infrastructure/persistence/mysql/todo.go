package mysql

import (
	"context"
	"database/sql"
	"github.com/sky0621/tips-go/experiment/architecture/onion_architecture/internal/domain"
	"github.com/sky0621/tips-go/experiment/architecture/onion_architecture/internal/infrastructure/db"
)

type todoRepository struct {
	q *db.Queries
}

func NewTodoRepository(dbConn *sql.DB) domain.TodoRepository {
	return &todoRepository{
		q: db.New(dbConn),
	}
}

func (r *todoRepository) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	id, err := r.q.InsertTodo(ctx, db.InsertTodoParams{
		Title:     todo.Title,
		Completed: todo.Completed,
	})
	if err != nil {
		return err
	}
	todo.ID = id
	return nil
}

func (r *todoRepository) GetTodos(ctx context.Context) ([]*domain.Todo, error) {
	rows, err := r.q.ListTodos(ctx)
	if err != nil {
		return nil, err
	}
	todos := make([]*domain.Todo, len(rows))
	for i, row := range rows {
		todos[i] = &domain.Todo{
			ID:        row.ID,
			Title:     row.Title,
			Completed: row.Completed,
		}
	}
	return todos, nil
}

func (r *todoRepository) GetTodoByID(ctx context.Context, id int64) (*domain.Todo, error) {
	row, err := r.q.GetTodo(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.Todo{
		ID:        row.ID,
		Title:     row.Title,
		Completed: row.Completed,
	}, nil
}

func (r *todoRepository) UpdateTodo(ctx context.Context, todo *domain.Todo) error {
	return r.q.UpdateTodo(ctx, db.UpdateTodoParams{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
	})
}

func (r *todoRepository) DeleteTodo(ctx context.Context, id int64) error {
	return r.q.DeleteTodo(ctx, id)
}
