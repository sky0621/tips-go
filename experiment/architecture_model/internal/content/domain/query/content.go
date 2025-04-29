package query

import (
	"context"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/domain/model/entity"
)

type Content interface {
	SearchContents(ctx context.Context, partialName *string) ([]entity.Content, error)
	GetContent(ctx context.Context, id string) (entity.Content, error)
}
