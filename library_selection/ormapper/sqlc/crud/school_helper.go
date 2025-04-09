package crud

import "github.com/sky0621/tips-go/library_selection/ormapper/sqlc/infra"

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

func aggregateRows(rows []infra.ListStudentsWithClassWithGradeWithSchoolRow) []School {
	schoolMap := make(map[int64]*School)

	for _, row := range rows {
		school, exists := schoolMap[row.SchoolID]
		if !exists {
			schoolMap[row.SchoolID] = newSchool(row)
			school = schoolMap[row.SchoolID]
		}

		gradeIndex := -1
		for i, grade := range school.Grades {
			if grade.GradeID == row.GradeID {
				gradeIndex = i
				break
			}
		}
		if gradeIndex == -1 {
			school.Grades = append(school.Grades, newGrade(row))
			gradeIndex = len(school.Grades) - 1
		}

		classIndex := -1
		for i, class := range school.Grades[gradeIndex].Classes {
			if class.ClassID == row.ClassID {
				classIndex = i
				break
			}
		}
		if classIndex == -1 {
			school.Grades[gradeIndex].Classes = append(school.Grades[gradeIndex].Classes, newClass(row))
			classIndex = len(school.Grades[gradeIndex].Classes) - 1
		}

		school.Grades[gradeIndex].Classes[classIndex].Students = append(
			school.Grades[gradeIndex].Classes[classIndex].Students,
			newStudent(row),
		)
	}

	// マップから最終結果のスライスに変換
	result := make([]School, 0, len(schoolMap))
	for _, s := range schoolMap {
		result = append(result, *s)
	}
	return result
}

func newStudent(row infra.ListStudentsWithClassWithGradeWithSchoolRow) Student {
	return Student{
		StudentID: row.StudentID,
		Name:      row.Name,
	}
}

func newClass(row infra.ListStudentsWithClassWithGradeWithSchoolRow) Class {
	return Class{
		ClassID:   row.ClassID,
		ClassName: row.ClassName,
		Students:  []Student{},
	}
}

func newGrade(row infra.ListStudentsWithClassWithGradeWithSchoolRow) Grade {
	return Grade{
		GradeID:   row.GradeID,
		GradeName: row.GradeName,
		Classes:   []Class{},
	}
}

func newSchool(row infra.ListStudentsWithClassWithGradeWithSchoolRow) *School {
	return &School{
		SchoolID:   row.SchoolID,
		SchoolName: row.SchoolName,
		Grades:     []Grade{},
	}
}

func extract[T any, R any](items []T, selector func(T) R) []R {
	result := make([]R, len(items))
	for i, item := range items {
		result[i] = selector(item)
	}
	return result
}
