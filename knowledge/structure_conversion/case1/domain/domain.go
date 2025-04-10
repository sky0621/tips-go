package domain

type School struct {
	SchoolID   int64
	SchoolName string

	Grades []Grade
}

type Grade struct {
	GradeID   int64
	GradeName string

	Classes []Class
}

type Class struct {
	ClassID   int64
	ClassName string

	Students []Student
}

type Student struct {
	StudentID int64
	Name      string
}
