package controller

import (
	"context"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/handler/interfaces"
)

type Course struct {
}

func (c Course) GetCourses(ctx context.Context, request interfaces.GetCoursesRequestObject) (interfaces.GetCoursesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c Course) PostCourses(ctx context.Context, request interfaces.PostCoursesRequestObject) (interfaces.PostCoursesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
