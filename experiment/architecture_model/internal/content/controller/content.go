package controller

import (
	"context"
	"github.com/google/uuid"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/application"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/domain/query"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/handler/interfaces"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/shared/converter"
)

func NewContent(q query.Content, saveService application.SaveContentService) Content {
	return Content{q: q, saveService: saveService}
}

type Content struct {
	q           query.Content
	saveService application.SaveContentService
}

func (c Content) PostContents(ctx context.Context, request interfaces.PostContentsRequestObject) (interfaces.PostContentsResponseObject, error) {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	if err := c.saveService.Save(ctx, uuidV7, request.Body.Name); err != nil {
		return nil, err
	}
	return interfaces.PostContents201JSONResponse(interfaces.ContentResponse{
		ID:   uuidV7.String(),
		Name: request.Body.Name,
	}), nil
}

func (c Content) GetContents(ctx context.Context, request interfaces.GetContentsRequestObject) (interfaces.GetContentsResponseObject, error) {
	contents, err := c.q.SearchContents(ctx, request.Params.PartialName)
	if err != nil {
		return nil, err
	}
	responses := make([]interfaces.ContentResponse, len(contents))
	for i, content := range contents {
		responses[i] = interfaces.ContentResponse{
			ID:   content.ID.String(),
			Name: content.Name,
		}
	}
	return interfaces.GetContents200JSONResponse(responses), nil
}

func (c Content) GetContentsByID(ctx context.Context, request interfaces.GetContentsByIDRequestObject) (interfaces.GetContentsByIDResponseObject, error) {
	_, err := uuid.Parse(request.ID)
	if err != nil {
		return interfaces.GetContentsByID400JSONResponse{Message: converter.ToPtr("not uuid")}, nil
	}
	content, err := c.q.GetContent(ctx, request.ID)
	if err != nil {
		return nil, err
	}
	if content.IsEmpty() {
		return interfaces.GetContentsByID404JSONResponse{Message: converter.ToPtr("not found")}, nil
	}
	return interfaces.GetContentsByID200JSONResponse(interfaces.ContentResponse{
		ID:   content.ID.String(),
		Name: content.Name,
	}), nil
}
