package main

import (
	"fmt"
	"github.com/sky0621/tips-go/knowledge/structure_conversion/case1/domain"
	"github.com/sky0621/tips-go/knowledge/structure_conversion/case1/infra"
	"github.com/sky0621/tips-go/knowledge/structure_conversion/case1/util"
)

func main() {
	iSchools := findSchools()
	iGrades := findGrades()
	iClasses := findClasses()
	iStudents := findStudents()

	dSchools := make([]domain.School, len(iSchools))
	for i, s := range iSchools {
		dSchools[i] = domain.School{
			SchoolID:   s.SchoolID,
			SchoolName: s.SchoolName,
		}
	}

	dSchoolMap := util.ToMap(dSchools, func(s domain.School) int64 {
		return s.SchoolID
	})
	fmt.Println("dSchoolMap:", dSchoolMap)

	// TODO:

	grades := util.ToMap(iGrades, func(g infra.Grade) int64 {
		return g.GradeID
	})
	fmt.Println("grades:", grades)

	classes := util.ToMap(iClasses, func(c infra.Class) int64 {
		return c.ClassID
	})
	fmt.Println("classes:", classes)

	students := util.ToMap(iStudents, func(s infra.Student) int64 {
		return s.StudentID
	})
	fmt.Println("students:", students)
}
