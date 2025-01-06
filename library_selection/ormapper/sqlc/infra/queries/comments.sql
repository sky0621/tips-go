-- name: CreateComment :execlastid
INSERT INTO comments (content, user_id, post_id) VALUES (?, ?, ?);

-- name: GetComment :one
SELECT id, content, user_id, post_id, created_at, updated_at
FROM comments
WHERE id = ? LIMIT 1;

-- name: DeleteComment :exec
DELETE FROM comments WHERE id = ?;

-- name: ListCommentsByPost :many
SELECT id, content, user_id, post_id, created_at, updated_at
FROM comments
WHERE post_id = ?;
