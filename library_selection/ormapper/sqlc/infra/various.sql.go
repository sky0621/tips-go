// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: various.sql

package infra

import (
	"context"
	"database/sql"
)

const listPostsByLikeTitle = `-- name: ListPostsByLikeTitle :many
SELECT id, title, content, user_id, created_at, updated_at FROM posts p WHERE p.title LIKE ?
`

func (q *Queries) ListPostsByLikeTitle(ctx context.Context, title string) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPostsByLikeTitle, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRecentCommentByPosts = `-- name: ListRecentCommentByPosts :many
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
    p.id
`

type ListRecentCommentByPostsRow struct {
	PostID                 int64
	PostTitle              string
	AuthorName             string
	LatestCommentContent   sql.NullString
	LatestCommenterName    sql.NullString
	LatestCommentCreatedAt sql.NullTime
}

func (q *Queries) ListRecentCommentByPosts(ctx context.Context) ([]ListRecentCommentByPostsRow, error) {
	rows, err := q.db.QueryContext(ctx, listRecentCommentByPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListRecentCommentByPostsRow{}
	for rows.Next() {
		var i ListRecentCommentByPostsRow
		if err := rows.Scan(
			&i.PostID,
			&i.PostTitle,
			&i.AuthorName,
			&i.LatestCommentContent,
			&i.LatestCommenterName,
			&i.LatestCommentCreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsersByIDs = `-- name: ListUsersByIDs :many
SELECT id, name, created_at, updated_at FROM users
WHERE id IN (?)
`

func (q *Queries) ListUsersByIDs(ctx context.Context, id int64) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsersByIDs, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsersWithPostAndCommentCount = `-- name: ListUsersWithPostAndCommentCount :many
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
    post_count DESC, comment_count DESC
`

type ListUsersWithPostAndCommentCountRow struct {
	UserID       int64
	UserName     string
	PostCount    int64
	CommentCount int64
}

func (q *Queries) ListUsersWithPostAndCommentCount(ctx context.Context) ([]ListUsersWithPostAndCommentCountRow, error) {
	rows, err := q.db.QueryContext(ctx, listUsersWithPostAndCommentCount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListUsersWithPostAndCommentCountRow{}
	for rows.Next() {
		var i ListUsersWithPostAndCommentCountRow
		if err := rows.Scan(
			&i.UserID,
			&i.UserName,
			&i.PostCount,
			&i.CommentCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsersWithRecentPostAndCommentCount = `-- name: ListUsersWithRecentPostAndCommentCount :many
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
    (recent_posts + recent_comments) DESC
`

type ListUsersWithRecentPostAndCommentCountRow struct {
	UserID         int64
	UserName       string
	RecentPosts    int64
	RecentComments int64
}

func (q *Queries) ListUsersWithRecentPostAndCommentCount(ctx context.Context) ([]ListUsersWithRecentPostAndCommentCountRow, error) {
	rows, err := q.db.QueryContext(ctx, listUsersWithRecentPostAndCommentCount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListUsersWithRecentPostAndCommentCountRow{}
	for rows.Next() {
		var i ListUsersWithRecentPostAndCommentCountRow
		if err := rows.Scan(
			&i.UserID,
			&i.UserName,
			&i.RecentPosts,
			&i.RecentComments,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const maxUsersID = `-- name: MaxUsersID :one
# https://docs.sqlc.dev/en/latest/howto/select.html#mysql-and-sqlite
# -- name: ListUsersByIDs :many
# SELECT * FROM users
# WHERE id IN (sqlc.slice('ids'));

SELECT MAX(id) AS maxId
FROM users
`

func (q *Queries) MaxUsersID(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, maxUsersID)
	var maxid interface{}
	err := row.Scan(&maxid)
	return maxid, err
}
