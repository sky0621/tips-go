package crud

import (
	"context"
	"log"
	"time"

	"github.com/sky0621/tips-go/library_selection/ormapper/bun/models"
	"github.com/uptrace/bun"
)

func Insert(ctx context.Context, db *bun.DB) {
	u := &models.User{
		Name:      "Alice",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := db.NewInsert().Model(u).Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	p := &models.Post{
		Title:     "First Post",
		Content:   "Hello Bun",
		UserID:    u.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = db.NewInsert().Model(p).Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	c := &models.Comment{
		Content:   "Nice post with Bun!",
		UserID:    u.ID,
		PostID:    p.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = db.NewInsert().Model(c).Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
