package crud

import (
	"context"
	"fmt"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/ent/ent"
)

func Select(ctx context.Context, client *ent.Client) {
	allUsers, err := client.User.Query().All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- 単一テーブル(users)の全件取得 --")
	fmt.Println(allUsers)

	allUsersWithPostsWithComments, err := client.User.Query().WithPosts().WithComments().All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- リレーションテーブル込みの全件取得 --")
	fmt.Println(allUsersWithPostsWithComments)
}
