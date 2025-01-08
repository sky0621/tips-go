package crud

import (
	"context"
	"fmt"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/infra"
)

func InsertBatch(ctx context.Context, q *infra.Queries) {
	usersBatch, err := q.CreateUsersBatch(ctx, infra.CreateUsersBatchParams{
		Name:   "S01",
		Name_2: "S02",
		Name_3: "S03",
		Name_4: "S04",
		Name_5: "S05",
	})
	if err != nil {
		log.Fatal("Insert batch error:", err)
	}
	fmt.Println("Insert batch success:", usersBatch)
	lastInsertId, err := usersBatch.LastInsertId()
	if err != nil {
		log.Fatal("Get last insert id error:", err)
	}
	fmt.Println("Get last insert id:", lastInsertId)
	rowsAffected, err := usersBatch.RowsAffected()
	if err != nil {
		log.Fatal("Get rowsAffected error:", err)
	}
	fmt.Println("Get rowsAffected count:", rowsAffected)
}
