package crud

import (
	"context"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/bun/models"
	"github.com/uptrace/bun"
)

func Delete(ctx context.Context, db *bun.DB) {
	_, err := db.NewDelete().Model(&models.User{ID: 1}).WherePK().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
