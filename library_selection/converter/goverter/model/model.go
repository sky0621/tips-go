package model

import "time"

type SchoolID string
type GradeID string
type ClassID string
type StudentID string

type School struct {
	ID        SchoolID
	Name      string
	No        int
	CreatedAt time.Time
	Grades    []Grade
}

type Grade struct {
	ID        GradeID
	GradeName string
	No        int
	CreatedAt time.Time
	Classes   []Class
}

type Class struct {
	ID        ClassID
	Name      string
	OnlyModel int
	Students  []Student
}

type Student struct {
	ID       SchoolID
	Name     string
	IsLeader bool
}
