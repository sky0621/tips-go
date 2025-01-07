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
	fmt.Println(rows)
}
