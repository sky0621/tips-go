package crud

import (
	"context"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/infra"
)

func Update(ctx context.Context, q *infra.Queries) {
	if err := q.UpdateUserName(ctx, infra.UpdateUserNameParams{Name: "Alice Updated", ID: 1}); err != nil {
		log.Fatal(err)
	}
}
