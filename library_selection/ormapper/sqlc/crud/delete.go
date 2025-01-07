package crud

import (
	"context"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/infra"
)

func Delete(ctx context.Context, q *infra.Queries) {
	if err := q.DeleteComment(ctx, 4); err != nil {
		log.Fatal(err)
	}
}
