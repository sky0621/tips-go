-- name: CreateContents :execlastid
INSERT INTO contents (id, name) VALUES (?, ?);

-- name: ListContents :many
SELECT id, name FROM contents;

-- name: SearchContents :many
SELECT id, name FROM contents WHERE name LIKE ?;
