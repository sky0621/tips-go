package query

import (
	"context"
	"github.com/google/uuid"
)

type SearchContents interface {
	Exec(ctx context.Context, partialName *string) ([]SearchContentsReadModel, error)
}

type SearchContentsReadModel struct {
	ID   uuid.UUID
	Name string
}
