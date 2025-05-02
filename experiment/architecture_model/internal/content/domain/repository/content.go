package repository

import (
	"context"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/domain/model"
)

type Content interface {
	Save(ctx context.Context, content model.ContentAggregate) error
}
