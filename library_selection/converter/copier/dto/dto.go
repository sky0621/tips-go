package dto

import (
	"fmt"
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
