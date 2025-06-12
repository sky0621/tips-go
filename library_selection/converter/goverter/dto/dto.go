package dto

import (
	"fmt"
	"github.com/sky0621/tips-go/library_selection/converter/goverter/model"
	"strings"
	"time"
)

type SchoolDTO struct {
	ID        string
	Name      string
	No        int
	CreatedAt time.Time
	Grades    []GradeDTO
}

type GradeDTO struct {
	ID        string
	Name      string
	No        int
	CreatedAt time.Time
	Classes   []ClassDTO
}

type ClassDTO struct {
	ID       string
	Name     string
	OnlyDTO  bool
	Students []StudentDTO
}

type StudentDTO struct {
	ID   string
	Name string
}

func (s SchoolDTO) String() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("School(ID: %s, Name: %s, No: %d, CreatedAt: %s)\n",
		s.ID, s.Name, s.No, s.CreatedAt.Format(time.RFC3339)))

	for _, g := range s.Grades {
		b.WriteString(fmt.Sprintf("  Grade(ID: %s, Name: %s, No: %d, CreatedAt: %s)\n",
			g.ID, g.Name, g.No, g.CreatedAt.Format(time.RFC3339)))

		for _, c := range g.Classes {
			b.WriteString(fmt.Sprintf("    Class(ID: %s, Name: %s)\n",
				c.ID, c.Name))

			for _, st := range c.Students {
				b.WriteString(fmt.Sprintf("      Student(ID: %s, Name: %s)\n",
					st.ID, st.Name))
			}
		}
	}

	return b.String()
}

// goverter:converter
// goverter:extend TimeToTime
type SchoolConverter interface {
	FromSchoolModel(source model.School) SchoolDTO
	// goverter:map GradeName Name
	FromGradeModel(source model.Grade) GradeDTO
	// goverter:map OnlyModel OnlyDTO | OnlyModelToOnlyDTO
	FromClassModel(source model.Class) ClassDTO

	ToSchoolModel(source SchoolDTO) model.School
	// goverter:map Name GradeName
	ToGradeModel(source GradeDTO) model.Grade
	// goverter:map OnlyDTO OnlyModel | OnlyDTOToOnlyModel
	ToClassModel(source ClassDTO) model.Class
	// goverter:ignore IsLeader
	ToStudentModel(source StudentDTO) model.Student
}

func TimeToTime(source time.Time) time.Time {
	return source
}

func OnlyModelToOnlyDTO(source int) bool {
	if source == 1 {
		return true
	}
	return false
}

func OnlyDTOToOnlyModel(source bool) int {
	if source {
		return 1
	}
	return 0
}
