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

	fmt.Println()
	fmt.Println("==========================")
	fmt.Println()

	employeesOrderBySalaryDesc, err := q.ListEmployeesOrderBySalaryDesc(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ListEmployeesOrderBySalaryDesc success")
	for _, s := range employeesOrderBySalaryDesc {
		fmt.Println(s)
	}

	fmt.Println()
	fmt.Println("==========================")
	fmt.Println()

	employeesOrderBySalaryAsc, err := q.ListEmployeesOrderBySalaryAsc(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ListEmployeesOrderBySalaryAsc success")
	for _, s := range employeesOrderBySalaryAsc {
		fmt.Println(s)
	}

	fmt.Println()
	fmt.Println("==========================")
	fmt.Println()

	employeesOrderByDepartmentIdDesc, err := q.ListEmployeesOrderByDepartmentIdDesc(ctx)
	if err != nil {
		return
	}
	fmt.Println("ListEmployeesOrderByDepartmentIdDesc success")
	for _, s := range employeesOrderByDepartmentIdDesc {
		fmt.Println(s)
	}

	fmt.Println()
	fmt.Println("==========================")
	fmt.Println()

	employeesOrderByDepartmentIdAsc, err := q.ListEmployeesOrderByDepartmentIdAsc(ctx)
	if err != nil {
		return
	}
	fmt.Println("ListEmployeesOrderByDepartmentIdAsc success")
	for _, s := range employeesOrderByDepartmentIdAsc {
		fmt.Println(s)
	}

	fmt.Println()
	fmt.Println("==========================")
	fmt.Println()

	employeesOrderByJoinDateDesc, err := q.ListEmployeesOrderByJoinDateDesc(ctx)
	if err != nil {
		return
	}
	fmt.Println("ListEmployeesOrderByJoinDateDesc success")
	for _, s := range employeesOrderByJoinDateDesc {
		fmt.Println(s)
	}

	fmt.Println()
	fmt.Println("==========================")
	fmt.Println()

	employeesOrderByJoinDateAsc, err := q.ListEmployeesOrderByJoinDateAsc(ctx)
	if err != nil {
		return
	}
	fmt.Println("ListEmployeesOrderByJoinDateAsc success")
	for _, s := range employeesOrderByJoinDateAsc {
		fmt.Println(s)
	}

}
