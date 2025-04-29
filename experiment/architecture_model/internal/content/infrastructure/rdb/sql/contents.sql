-- name: CreateContents :exec
INSERT INTO contents (id, name) VALUES (?, ?);

-- name: ListContents :many
SELECT id, name FROM contents;

-- name: SearchContents :many
SELECT id, name FROM contents WHERE name LIKE ?;

-- name: GetContentById :one
SELECT id, name FROM contents WHERE id = UUID_TO_BIN(?);
