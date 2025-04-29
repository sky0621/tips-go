package query

import (
	"context"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/domain/model"
)

type Content interface {
	SearchContents(ctx context.Context, partialName *string) ([]model.Content, error)
}
