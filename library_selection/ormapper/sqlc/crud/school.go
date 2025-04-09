package crud

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/infra"
	"log"
)

func SchoolSetup(ctx context.Context, q *infra.Queries) {
	_, err := q.CreateSchool(ctx)
	if err != nil {
		log.Fatal(err)
	}
	_, err = q.CreateGrade(ctx)
	if err != nil {
		log.Fatal(err)
	}
	_, err = q.CreateClassBatch(ctx)
	if err != nil {
		log.Fatal(err)
	}
	_, err = q.CreateStudentBatch(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func GetSchool(ctx context.Context, q *infra.Queries) {
	students, err := q.ListStudentsWithClassWithGradeWithSchool(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	result := aggregateRows(students)
	fmt.Println(result)
}

func GetSchool2(ctx context.Context, q *infra.Queries) {
	school, err := q.GetSchoolByID(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	grades, err := q.ListGradeBySchoolID(ctx, sql.NullInt64{Int64: school.SchoolID, Valid: true})
	if err != nil {
		return
	}
	for _, grade := range grades {
		classes, err := q.ListClassByGradeID(ctx, sql.NullInt64{Int64: grade.GradeID, Valid: true})
		if err != nil {
			return
		}
		for _, class := range classes {
			students, err := q.ListStudentByClassID(ctx, sql.NullInt64{Int64: class.ClassID, Valid: true})
			if err != nil {
				return
			}
			for _, student := range students {
				fmt.Println(student)
			}
		}
	}
}

func GetSchool3(ctx context.Context, q *infra.Queries) {
	school, err := q.GetSchoolByID(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	grades, err := q.ListGradeBySchoolID(ctx, sql.NullInt64{Int64: school.SchoolID, Valid: true})
	if err != nil {
		return
	}

	gradeIDs := extract(grades, func(g infra.Grade) sql.NullInt64 { return sql.NullInt64{Int64: g.GradeID, Valid: true} })
	classes, err := q.ListClassInGradeIDs(ctx, gradeIDs)

	classIDs := extract(classes, func(c infra.Class) sql.NullInt64 { return sql.NullInt64{Int64: c.ClassID, Valid: true} })
	students, err := q.ListStudentInClassIDs(ctx, classIDs)

	fmt.Println(students)
}
