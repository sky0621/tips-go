package migration

import (
	"context"
	"github.com/sky0621/tips-go/library_selection/ormapper/bun/models"
	"github.com/uptrace/bun"
	"log"
)

func Migrate(db *bun.DB) {
	ctx := context.Background()
	err := db.ResetModel(ctx,
		&models.User{},
		&models.Post{},
		&models.Comment{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
