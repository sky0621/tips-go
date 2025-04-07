package crud

import (
	"context"
	"fmt"
	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/infra"
	"log"
)

func Search(ctx context.Context, q *infra.Queries) {
	condition1 := &infra.EmployeeSearchCondition{EmployeeID: 104, FirstName: "Bob"}
	result1, err := q.SearchEmployeesWithCondition(ctx, condition1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SearchEmployeesWithCondition: {EmployeeID: 104, FirstName: \"Bob\"}")
	for _, s := range result1 {
		fmt.Println(s)
	}
}
