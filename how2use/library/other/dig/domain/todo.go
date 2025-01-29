package domain

type Todo struct {
	ID      int
	Content string
}

type TodoRepository interface {
	GetTodo(id int) (*Todo, error)
	AddTodo(content string) (*Todo, error)
}
