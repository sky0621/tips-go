package crud

import (
	"context"
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
