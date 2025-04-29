package controller

import (
	"context"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/domain/query"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/handler/interfaces"
)

type Content struct {
	q query.Content
}

func (c Content) PostContents(ctx context.Context, request interfaces.PostContentsRequestObject) (interfaces.PostContentsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c Content) GetContents(ctx context.Context, request interfaces.GetContentsRequestObject) (interfaces.GetContentsResponseObject, error) {
	contents, err := c.q.SearchContents(ctx, request.Params.PartialName)
	if err != nil {
		return nil, err
	}
	responses := make([]interfaces.ContentResponse, len(contents))
	for _, content := range contents {
		responses = append(responses, interfaces.ContentResponse{
			ID:   content.ID,
			Name: content.Name,
		})
	}
	return interfaces.GetContents200JSONResponse(responses), nil
}
