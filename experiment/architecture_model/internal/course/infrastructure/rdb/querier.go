// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package rdb

import (
	"context"
)

type Querier interface {
	CreateCourses(ctx context.Context, arg CreateCoursesParams) (int64, error)
	ListCourses(ctx context.Context) ([]Course, error)
	SearchCourses(ctx context.Context, level int32) ([]Course, error)
}

var _ Querier = (*Queries)(nil)
