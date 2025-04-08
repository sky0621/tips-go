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

func School(ctx context.Context, q *infra.Queries) {
	students, err := q.ListStudentsWithClassWithGradeWithSchool(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	for _, student := range students {
		fmt.Println(student)
	}
}

func School2(ctx context.Context, q *infra.Queries) {
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
