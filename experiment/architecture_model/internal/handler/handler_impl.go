package handler

import (
	"context"
	contentsController "github.com/sky0621/tips-go/experiment/architecture_model/internal/content/controller"
	coursesController "github.com/sky0621/tips-go/experiment/architecture_model/internal/course/controller"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/handler/interfaces"
)

var _ interfaces.StrictServerInterface = (*strictServerImpl)(nil)

func New(
	contentsController contentsController.Content,
	coursesController coursesController.Course,
	middlewares []interfaces.StrictMiddlewareFunc,
) interfaces.ServerInterface {
	return interfaces.NewStrictHandler(&strictServerImpl{
		contentsController,
		coursesController,
	}, middlewares)
}

type strictServerImpl struct {
	contentsController contentsController.Content
	coursesController  coursesController.Course
}

func (s strictServerImpl) GetContents(ctx context.Context, request interfaces.GetContentsRequestObject) (interfaces.GetContentsResponseObject, error) {
	return s.contentsController.GetContents(ctx, request)
}

func (s strictServerImpl) GetContentsByID(ctx context.Context, request interfaces.GetContentsByIDRequestObject) (interfaces.GetContentsByIDResponseObject, error) {
	return s.contentsController.GetContentsByID(ctx, request)
}

func (s strictServerImpl) PostContents(ctx context.Context, request interfaces.PostContentsRequestObject) (interfaces.PostContentsResponseObject, error) {
	return s.contentsController.PostContents(ctx, request)
}

func (s strictServerImpl) GetCourses(ctx context.Context, request interfaces.GetCoursesRequestObject) (interfaces.GetCoursesResponseObject, error) {
	return s.coursesController.GetCourses(ctx, request)
}

func (s strictServerImpl) PostCourses(ctx context.Context, request interfaces.PostCoursesRequestObject) (interfaces.PostCoursesResponseObject, error) {
	return s.coursesController.PostCourses(ctx, request)
}
