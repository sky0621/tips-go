package crud

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/infra"
)

func Select(ctx context.Context, q *infra.Queries) {
	user, err := q.GetUser(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- 単一テーブル(users)の１件取得 --")
	fmt.Println(user)

	allUsers, err := q.ListUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- 単一テーブル(users)の全件取得 --")
	fmt.Println(allUsers)

	postsByUser, err := q.ListPostsByUser(ctx, sql.NullInt64{
		Int64: 1,
		Valid: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- userに紐づくposts取得 --")
	fmt.Println(postsByUser)

	userWithPostAndComments, err := q.ListUserWithPostAndComments(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- 紐づくpostsとcommentsも含めてusers取得 --")
	fmt.Println(userWithPostAndComments)
}
