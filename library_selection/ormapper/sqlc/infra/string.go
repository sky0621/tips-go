package infra

import (
	"encoding/json"
	"fmt"
	"time"
)

// Department の String() メソッド
func (d Department) String() string {
	data, err := json.Marshal(d)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return string(data)
}

// Employee の String() メソッド
func (e Employee) String() string {
	// 匿名構造体を用いて、Null系フィールドは有効なら値、無効なら nil を設定する
	emp := struct {
		EmployeeID   int64      `json:"EmployeeID"`
		FirstName    string     `json:"FirstName"`
		LastName     string     `json:"LastName"`
		Salary       *string    `json:"Salary"`
		DepartmentID *int64     `json:"DepartmentID"`
		JoinDate     *time.Time `json:"JoinDate"`
	}{
		EmployeeID: e.EmployeeID,
		FirstName:  e.FirstName,
		LastName:   e.LastName,
	}

	if e.Salary.Valid {
		emp.Salary = &e.Salary.String
	}
	if e.DepartmentID.Valid {
		emp.DepartmentID = &e.DepartmentID.Int64
	}
	if e.JoinDate.Valid {
		emp.JoinDate = &e.JoinDate.Time
	}

	data, err := json.Marshal(emp)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return string(data)
}
