package crud

import (
	"context"
	"fmt"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/infra"
)

func Various(ctx context.Context, q *infra.Queries) {
	rows, err := q.ListUsersWithPostAndCommentCount(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- 紐づくpostsとcommentsの件数を含めてusers取得 --")
	for _, row := range rows {
		fmt.Println(row)
	}

	fmt.Println()

	rows2, err := q.ListUsersWithRecentPostAndCommentCount(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- 紐づく最近のpostsとcommentsの件数を含めてusers取得 --")
	for _, row2 := range rows2 {
		fmt.Println(row2)
	}

	fmt.Println()

	rows3, err := q.ListRecentCommentByPosts(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- 投稿ごとの最新コメント取得 --")
	for _, row3 := range rows3 {
		fmt.Println(row3)
	}
}
