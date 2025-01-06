-- name: CreatePost :execlastid
INSERT INTO posts (title, content, user_id) VALUES (?, ?, ?);

-- name: GetPost :one
SELECT id, title, content, user_id, created_at, updated_at
FROM posts
WHERE id = ? LIMIT 1;

-- name: UpdatePost :exec
UPDATE posts
SET title = ?, content = ?
WHERE id = ?;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = ?;

-- name: ListPostsByUser :many
SELECT id, title, content, user_id, created_at, updated_at
FROM posts
WHERE user_id = ?;
