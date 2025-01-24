package crud

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/infra"
)

func Complex(ctx context.Context, q *infra.Queries) {
	query, err := q.ListWithComplexQuery(ctx, infra.ListWithComplexQueryParams{
		UserID:   sql.NullInt64{Int64: 1, Valid: true},
		UserID_2: sql.NullInt64{Int64: 1, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, q := range query {
		fmt.Println(q)
	}
}
