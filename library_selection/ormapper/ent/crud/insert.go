package crud

import (
	"context"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/ent/ent"
)

func Insert(ctx context.Context, client *ent.Client) {
	user, err := client.User.Create().
		SetName("Alice").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed inserting record: %v", err)
	}

	post, err := client.Post.Create().
		SetTitle("First Post").
		SetContent("Hello ent").
		SetUsers(user).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed inserting record: %v", err)
	}

	_, err = client.Comment.Create().
		SetContent("Nice post ent!").
		SetUsers(user).
		SetPosts(post).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed inserting record: %v", err)
	}
}
