-- name: ListUsersWithPostAndCommentCount :many
SELECT
    users.id AS user_id,
    users.name AS user_name,
    COUNT(DISTINCT posts.id) AS post_count,
    COUNT(comments.id) AS comment_count
FROM
    users
        LEFT JOIN
    posts ON users.id = posts.user_id
        LEFT JOIN
    comments ON users.id = comments.user_id AND posts.id = comments.post_id
GROUP BY
    users.id, users.name
ORDER BY
    post_count DESC, comment_count DESC;

