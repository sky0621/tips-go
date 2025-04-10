package infra

import "database/sql"

type School struct {
	SchoolID   int64
	SchoolName string
}

type Grade struct {
	GradeID   int64
	GradeName string
	SchoolID  sql.NullInt64
}

type Class struct {
	ClassID   int64
	ClassName string
	GradeID   sql.NullInt64
}

type Student struct {
	StudentID int64
	Name      string
	ClassID   sql.NullInt64
}
