package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/application/query"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/infrastructure/rdb"
)

var _ query.GetContent = (*getContentImpl)(nil)

func NewGetContent(db *sql.DB) query.GetContent {
	return &getContentImpl{db: db}
}

type getContentImpl struct {
	db *sql.DB
}

func (g getContentImpl) Exec(ctx context.Context, id string) (query.GetContentReadModel, error) {
	q := rdb.New(g.db)
	content, err := q.GetContentById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return query.GetContentReadModel{}, nil
		}
		return query.GetContentReadModel{}, err
	}
	uuidID, err := uuid.FromBytes(content.ID)
	if err != nil {
		return query.GetContentReadModel{}, err
	}
	return query.GetContentReadModel{
		ID:   uuidID,
		Name: content.Name,
	}, nil
}
