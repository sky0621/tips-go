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

-- name: ListUsersWithRecentPostAndCommentCount :many
SELECT
    users.id AS user_id,
    users.name AS user_name,
    COUNT(DISTINCT posts.id) AS recent_posts,
    COUNT(DISTINCT comments.id) AS recent_comments
FROM
    users
        LEFT JOIN
    posts ON users.id = posts.user_id AND posts.created_at >= NOW() - INTERVAL 1 MONTH
    LEFT JOIN
    comments ON users.id = comments.user_id AND comments.created_at >= NOW() - INTERVAL 1 MONTH
GROUP BY
    users.id, users.name
HAVING
    recent_posts > 0 OR recent_comments > 0
ORDER BY
    (recent_posts + recent_comments) DESC;

-- name: ListRecentCommentByPosts :many
SELECT
    p.id AS post_id,
    p.title AS post_title,
    u.name AS author_name,
    c.content AS latest_comment_content,
    cu.name AS latest_commenter_name,
    c.created_at AS latest_comment_created_at
FROM
    posts p
        JOIN
    users u ON p.user_id = u.id
        LEFT JOIN
    comments c ON p.id = c.post_id
        LEFT JOIN
    users cu ON c.user_id = cu.id
WHERE
    c.created_at = (
        SELECT MAX(created_at)
        FROM comments
        WHERE post_id = p.id
    )
ORDER BY
    p.id;
