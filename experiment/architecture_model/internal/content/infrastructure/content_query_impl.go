package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/domain/model/entity"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/domain/query"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/infrastructure/rdb"
)

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

func (c contentQueryImpl) GetContent(ctx context.Context, id string) (entity.Content, error) {
	q := rdb.New(c.db)
	content, err := q.GetContentById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Content{}, nil
		}
		return entity.Content{}, err
	}
	uuidID, err := uuid.FromBytes(content.ID)
	if err != nil {
		return entity.Content{}, err
	}
	return entity.Content{
		ID:   uuidID,
		Name: content.Name,
	}, nil
}
