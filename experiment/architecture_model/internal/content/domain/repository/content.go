package repository

import (
	"context"
	"github.com/google/uuid"
)

type Content interface {
	Save(ctx context.Context, id uuid.UUID, name string) error
}
