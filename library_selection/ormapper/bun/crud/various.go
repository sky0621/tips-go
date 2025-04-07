package crud

import (
	"context"
	"fmt"
	"github.com/sky0621/tips-go/library_selection/ormapper/bun/models"
	"github.com/uptrace/bun"
	"log"
)

func Various(ctx context.Context, db *bun.DB) {
	var users []models.User
	if err := db.NewSelect().Model(&models.User{}).
		Where("id = ?", 1).
		Where("name = ?", "Alice").
		Scan(ctx, &users); err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Println(user)
	}
}
