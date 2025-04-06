package crud

import (
	"context"
	"fmt"
	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/infra"
	"log"
)

func OtherSetup(ctx context.Context, q *infra.Queries) {
	departmentsBatch, err := q.CreateDepartmentsBatch(ctx)
	if err != nil {
		log.Fatal("Insert batch error:", err)
	}
	fmt.Println("Insert batch success:", departmentsBatch)

	employeesBatch, err := q.CreateEmployeesBatch(ctx)
	if err != nil {
		log.Fatal("Insert batch error:", err)
	}
	fmt.Println("Insert batch success:", employeesBatch)
}

func Other(ctx context.Context, q *infra.Queries) {
	salaries, err := q.ListEmployeesOrderBySalary(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ListEmployeesOrderBySalary success")
	for _, s := range salaries {
		fmt.Println(s)
	}

	fmt.Println()
	fmt.Println("==========================")
	fmt.Println()

	/*
	 * ここでカラム名を指定しても、並べ替えは機能しない。。
	 */
	xxxx, err := q.ListEmployeesOrderByXXXX(ctx, "salary")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ListEmployeesOrderByXXXX success")
	for _, s := range xxxx {
		fmt.Println(s)
	}

}
