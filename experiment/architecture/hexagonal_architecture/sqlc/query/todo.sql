-- name: InsertTodo :execlastid
INSERT INTO todos (title, completed) VALUES (?, ?);

-- name: ListTodos :many
SELECT id, title, completed FROM todos;

-- name: GetTodo :one
SELECT id, title, completed FROM todos WHERE id = ?;

-- name: UpdateTodo :exec
UPDATE todos SET title = ?, completed = ? WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = ?;
