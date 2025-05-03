package controller

import (
	"context"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/api"
)

type Course struct {
}

func (c Course) GetCourses(ctx context.Context, request api.GetCoursesRequestObject) (api.GetCoursesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c Course) PostCourses(ctx context.Context, request api.PostCoursesRequestObject) (api.PostCoursesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
