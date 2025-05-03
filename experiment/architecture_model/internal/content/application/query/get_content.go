package query

import (
	"context"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/shared/model"
)

type GetContent interface {
	Exec(ctx context.Context, id string) (GetContentReadModel, error)
}

type GetContentReadModel struct {
	ID   model.ID
	Name string
}

func (m GetContentReadModel) IsEmpty() bool {
	return m.ID.IsEmpty() && m.Name == ""
}
