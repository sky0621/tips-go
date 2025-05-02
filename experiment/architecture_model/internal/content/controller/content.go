package controller

import (
	"context"
	"github.com/google/uuid"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/application/command"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/application/query"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/domain/model"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/handler/interfaces"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/shared/converter"
)

func NewContent(searchContents query.SearchContents, getContent query.GetContent, saveContent command.SaveContent) Content {
	return Content{searchContents: searchContents, getContent: getContent, saveContent: saveContent}
}

type Content struct {
	searchContents query.SearchContents
	getContent     query.GetContent
	saveContent    command.SaveContent
}

func (c Content) PostContents(ctx context.Context, request interfaces.PostContentsRequestObject) (interfaces.PostContentsResponseObject, error) {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	programs := make([]model.ProgramWriteModel, len(request.Body.Programs))
	for i, program := range request.Body.Programs {
		uuidV7, err := uuid.NewV7()
		if err != nil {
			return nil, err
		}
		programs[i] = model.ProgramWriteModel{
			ID:       uuidV7,
			Question: program.Question,
			Answer:   program.Answer,
		}
	}

	if err := c.saveContent.Save(ctx, model.SaveContentWriteModel{
		ID:       uuidV7,
		Name:     request.Body.Name,
		Programs: programs,
	}); err != nil {
		return nil, err
	}
	return interfaces.PostContents201JSONResponse(interfaces.ContentResponse{
		ID:   uuidV7.String(),
		Name: request.Body.Name,
	}), nil
}

func (c Content) GetContents(ctx context.Context, request interfaces.GetContentsRequestObject) (interfaces.GetContentsResponseObject, error) {
	contents, err := c.searchContents.Exec(ctx, request.Params.PartialName)
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
	content, err := c.getContent.Exec(ctx, request.ID)
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
