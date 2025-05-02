package infrastructure

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/application/query"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/infrastructure/rdb"
)

var _ query.SearchContents = (*searchContentImpl)(nil)

func NewSearchContents(db *sql.DB) query.SearchContents {
	return &searchContentImpl{db: db}
}

type searchContentImpl struct {
	db *sql.DB
}

func (s searchContentImpl) Exec(ctx context.Context, partialName *string) ([]query.SearchContentsReadModel, error) {
	q := rdb.New(s.db)
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
	results := make([]query.SearchContentsReadModel, len(contents))
	for i, content := range contents {
		id, err := uuid.FromBytes(content.ID)
		if err != nil {
			return nil, err
		}
		results[i] = query.SearchContentsReadModel{
			ID:   id,
			Name: content.Name,
		}
	}
	return results, nil
}
