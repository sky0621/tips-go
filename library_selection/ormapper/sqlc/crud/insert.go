package crud

import (
	"context"
	"database/sql"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/infra"
)

func Insert(ctx context.Context, q *infra.Queries) {
	userID, err := q.CreateUser(ctx, "Alice")
	if err != nil {
		log.Fatal(err)
	}

	var postID int64
	postID, err = q.CreatePost(ctx, infra.CreatePostParams{
		Title:   "First Post",
		Content: sql.NullString{String: "Hello sqlc", Valid: true},
		UserID:  sql.NullInt64{Int64: userID, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	var postID2 int64
	postID2, err = q.CreatePost(ctx, infra.CreatePostParams{
		Title:   "Second Post",
		Content: sql.NullString{String: "Hello sqlc2", Valid: true},
		UserID:  sql.NullInt64{Int64: userID, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = q.CreateComment(ctx, infra.CreateCommentParams{
		Content: sql.NullString{String: "Nice post with sqlc!", Valid: true},
		UserID:  sql.NullInt64{Int64: userID, Valid: true},
		PostID:  sql.NullInt64{Int64: postID, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	_, err = q.CreateComment(ctx, infra.CreateCommentParams{
		Content: sql.NullString{String: "Nice post with sqlc2!", Valid: true},
		UserID:  sql.NullInt64{Int64: userID, Valid: true},
		PostID:  sql.NullInt64{Int64: postID, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	_, err = q.CreateComment(ctx, infra.CreateCommentParams{
		Content: sql.NullString{String: "Nice post with sqlc3!", Valid: true},
		UserID:  sql.NullInt64{Int64: userID, Valid: true},
		PostID:  sql.NullInt64{Int64: postID2, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	_, err = q.CreateComment(ctx, infra.CreateCommentParams{
		Content: sql.NullString{String: "Nice post with sqlc4!", Valid: true},
		UserID:  sql.NullInt64{Int64: userID, Valid: true},
		PostID:  sql.NullInt64{Int64: postID2, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}
}
