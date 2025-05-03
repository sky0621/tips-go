package query

import (
	"context"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/shared/model"
)

type SearchContents interface {
	Exec(ctx context.Context, partialName *string) ([]SearchContentsReadModel, error)
}

type SearchContentsReadModel struct {
	ID   model.ID
	Name string
}
