package infrastructure

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/domain/model/entity"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/domain/query"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/infrastructure/rdb"
)

var _ query.Content = (*contentQueryImpl)(nil)

func NewContentQuery(db *sql.DB) query.Content {
	return &contentQueryImpl{db: db}
}

type contentQueryImpl struct {
	db *sql.DB
}

func (c contentQueryImpl) SearchContents(ctx context.Context, partialName *string) ([]entity.Content, error) {
	q := rdb.New(c.db)
	var contents []rdb.Content
	var err error
	if partialName == nil {
		contents, err = q.ListContents(ctx)
	} else {
		contents, err = q.SearchContents(ctx, *partialName)
	}
	if err != nil {
		return nil, err
	}
	results := make([]entity.Content, len(contents))
	for i, content := range contents {
		id, err := uuid.FromBytes(content.ID)
		if err != nil {
			return nil, err
		}
		results[i] = entity.Content{
			ID:   id,
			Name: content.Name,
		}
	}
	return results, nil
}
