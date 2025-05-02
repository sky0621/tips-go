package query

import (
	"context"
	"github.com/google/uuid"
)

type GetContent interface {
	Exec(ctx context.Context, id string) (GetContentReadModel, error)
}

type GetContentReadModel struct {
	ID   uuid.UUID
	Name string
}

func (m GetContentReadModel) IsEmpty() bool {
	return m.ID == uuid.Nil && m.Name == ""
}
