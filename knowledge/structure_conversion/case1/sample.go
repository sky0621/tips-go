package main

import (
	"database/sql"
	"github.com/sky0621/tips-go/knowledge/structure_conversion/case1/infra"
)

func findSchools() []infra.School {
	return []infra.School{
		{SchoolID: 1, SchoolName: "学校A"},
		{SchoolID: 2, SchoolName: "学校B"},
	}
}

func findGrades() []infra.Grade {
	return []infra.Grade{
		{GradeID: 1, GradeName: "１学年", SchoolID: sql.NullInt64{Int64: 1, Valid: true}},
		{GradeID: 2, GradeName: "２学年", SchoolID: sql.NullInt64{Int64: 1, Valid: true}},
		{GradeID: 3, GradeName: "１学年", SchoolID: sql.NullInt64{Int64: 2, Valid: true}},
	}
}

func findClasses() []infra.Class {
	return []infra.Class{
		{ClassID: 1, ClassName: "１−１", GradeID: sql.NullInt64{Int64: 1, Valid: true}},
		{ClassID: 2, ClassName: "１−２", GradeID: sql.NullInt64{Int64: 1, Valid: true}},
		{ClassID: 3, ClassName: "２−１", GradeID: sql.NullInt64{Int64: 2, Valid: true}},
		{ClassID: 4, ClassName: "１−１", GradeID: sql.NullInt64{Int64: 3, Valid: true}},
		{ClassID: 5, ClassName: "１−２", GradeID: sql.NullInt64{Int64: 3, Valid: true}},
	}
}

func findStudents() []infra.Student {
	return []infra.Student{
		{StudentID: 1, Name: "山田太郎", ClassID: sql.NullInt64{Int64: 1, Valid: true}},
		{StudentID: 2, Name: "山田花子", ClassID: sql.NullInt64{Int64: 1, Valid: true}},
		{StudentID: 3, Name: "鈴木一郎", ClassID: sql.NullInt64{Int64: 2, Valid: true}},
		{StudentID: 4, Name: "佐藤二郎", ClassID: sql.NullInt64{Int64: 2, Valid: true}},
		{StudentID: 5, Name: "田中三郎", ClassID: sql.NullInt64{Int64: 2, Valid: true}},
		{StudentID: 6, Name: "山田四郎", ClassID: sql.NullInt64{Int64: 3, Valid: true}},
		{StudentID: 7, Name: "山田五郎", ClassID: sql.NullInt64{Int64: 4, Valid: true}},
		{StudentID: 8, Name: "山田六郎", ClassID: sql.NullInt64{Int64: 4, Valid: true}},
		{StudentID: 9, Name: "山田七郎", ClassID: sql.NullInt64{Int64: 4, Valid: true}},
		{StudentID: 10, Name: "山田八郎", ClassID: sql.NullInt64{Int64: 4, Valid: true}},
		{StudentID: 11, Name: "山田九郎", ClassID: sql.NullInt64{Int64: 5, Valid: true}},
	}
}
