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

	fmt.Println()

	posts, err := q.ListPostsByLikeTitle(ctx, "%Post")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- title指定でposts取得 --")
	for _, post := range posts {
		fmt.Printf("ID: %u, Title: %s\n", post.ID, post.Title)
	}

	fmt.Println()

	users, err := q.ListUsersByIDs(ctx, []int64{1, 2, 3, 4, 5, 6, 99})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- id指定でusers取得 --")
	for _, u := range users {
		fmt.Printf("ID: %d, Name: %s\n", u.ID, u.Name)
	}

	fmt.Println()

	relations, err := q.Relations(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- relations --")
	for _, r := range relations {
		fmt.Printf("ID: %d, Name: %s\n", r.ID, r.Name)
	}

	fmt.Println()

	id, err := q.MaxUsersID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- usersからMAX(id)を取得 --")
	fmt.Println(id)
}
