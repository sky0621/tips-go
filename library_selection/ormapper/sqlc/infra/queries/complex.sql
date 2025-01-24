-- name: ListUserWithPostAndComments :many
SELECT
    users.id AS user_id,
    users.name AS user_name,
    users.created_at AS user_created_at,
    users.updated_at AS user_updated_at,
    posts.id AS post_id,
    posts.title AS post_title,
    posts.content AS post_content,
    posts.created_at AS post_created_at,
    posts.updated_at AS post_updated_at,
    comments.id AS comment_id,
    comments.content AS comment_content,
    comments.created_at AS comment_created_at,
    comments.updated_at AS comment_updated_at
FROM users
INNER JOIN posts ON posts.user_id = users.id
INNER JOIN comments ON comments.user_id = users.id AND comments.post_id = posts.id;

-- name: ListWithComplexQuery :many
WITH
    c AS (SELECT * FROM comments WHERE comments.user_id = ?),
    p AS (SELECT * FROM posts WHERE posts.user_id = ?)
SELECT
    c.user_id AS user_id,
    p.id AS post_id,
    c.content AS comment_content,
    p.content AS post_content,
    p.title AS post_title
FROM c INNER JOIN p ON p.user_id = c.user_id;
