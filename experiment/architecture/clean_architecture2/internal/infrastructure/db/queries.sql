-- name: GetUserByID :one
SELECT id, name FROM users WHERE id = ?;
