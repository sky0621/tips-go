-- name: CreateDepartmentsBatch :execresult
INSERT INTO departments (department_id, department_name)
VALUES (1, 'Sales'),
       (2, 'Engineering'),
       (3, 'HR'),
       (4, 'Finance');

-- name: CreateEmployeesBatch :execresult
INSERT INTO employees (employee_id, first_name, last_name, salary, department_id, join_date)
VALUES (101, 'John', 'Doe', 55000.00, 1, '2020-01-15'),
       (102, 'Jane', 'Smith', 75000.00, 2, '2019-06-01'),
       (103, 'Alice', 'Brown', 80000.00, 2, '2021-07-20'),
       (104, 'Bob', 'Davis', 50000.00, 3, '2018-03-10'),
       (105, 'Charlie', 'Evans', 65000.00, 1, '2022-11-05'),
       (106, 'Eve', 'White', 90000.00, 4, '2017-05-23'),
       (107, 'Frank', 'Green', 60000.00, 3, '2023-01-10'),
       (108, 'Grace', 'Hall', 70000.00, 4, '2019-09-15');

# これは機能しない。
-- name: ListEmployeesOrderByXXXX :many
SELECT * FROM employees ORDER BY ? DESC;

-- name: ListEmployeesOrderBySalaryDesc :many
SELECT * FROM employees ORDER BY salary DESC;
-- name: ListEmployeesOrderBySalaryAsc :many
SELECT * FROM employees ORDER BY salary;

-- name: ListEmployeesOrderByDepartmentIdDesc :many
SELECT * FROM employees ORDER BY department_id DESC;
-- name: ListEmployeesOrderByDepartmentIdAsc :many
SELECT * FROM employees ORDER BY department_id;

-- name: ListEmployeesOrderByJoinDateDesc :many
SELECT * FROM employees ORDER BY join_date DESC;
-- name: ListEmployeesOrderByJoinDateAsc :many
SELECT * FROM employees ORDER BY join_date;
