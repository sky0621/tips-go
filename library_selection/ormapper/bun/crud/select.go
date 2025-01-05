package crud

import (
	"context"
	"fmt"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/bun/models"

	"github.com/uptrace/bun"
)

func Select(ctx context.Context, db *bun.DB) {
	var allUsers []models.User
	if err := db.NewSelect().Model(&allUsers).Scan(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- 単一テーブル(users)の全件取得 --")
	fmt.Println(allUsers)

	var allUsersWithPostsWithComments []models.User
	if err := db.NewSelect().Model(&allUsersWithPostsWithComments).Relation("Posts").Scan(ctx); err != nil {
	}
	fmt.Println("-- リレーションテーブル込みの全件取得 --")
	fmt.Println(allUsersWithPostsWithComments)
}
