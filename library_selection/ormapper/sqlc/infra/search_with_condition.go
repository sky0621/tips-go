package infra

import (
	"context"
	"log"
	"strings"
)

const searchEmployeesWithCondition = searchEmployees + " WHERE "

type EmployeeSearchCondition struct {
	EmployeeID int64
	FirstName  string
	LastName   string
}

func (q *Queries) SearchEmployeesWithCondition(ctx context.Context, condition *EmployeeSearchCondition) ([]Employee, error) {
	if condition == nil {
		return q.SearchEmployees(ctx)
	}
	conditions := make([]string, 0)
	args := make([]any, 0)
	if condition.FirstName != "" {
		conditions = append(conditions, "first_name = ?")
		args = append(args, condition.FirstName)
	}
	if condition.LastName != "" {
		conditions = append(conditions, "last_name = ?")
		args = append(args, condition.LastName)
	}
	if condition.EmployeeID != 0 {
		conditions = append(conditions, "employee_id = ?")
		args = append(args, condition.EmployeeID)
	}
	query := searchEmployeesWithCondition + strings.Join(conditions, " AND ")
	log.Println(query)
	rows, err := q.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Employee{}
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.EmployeeID,
			&i.FirstName,
			&i.LastName,
			&i.Salary,
			&i.DepartmentID,
			&i.JoinDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
