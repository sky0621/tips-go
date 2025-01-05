package crud

import (
	"context"
	"log"
	"time"

	"github.com/sky0621/tips-go/library_selection/ormapper/bun/models"
	"github.com/uptrace/bun"
)

func Update(ctx context.Context, db *bun.DB) {
	var u models.User
	if err := db.NewSelect().Model(&u).Where("id = ?", 1).Scan(ctx, &u); err != nil {
		log.Fatal(err)
	}
	u.Name = "Alice Updated"
	u.UpdatedAt = time.Now()
	_, err := db.NewUpdate().Model(&u).WherePK().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
