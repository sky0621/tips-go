package infrastructure

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/domain/repository"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/infrastructure/rdb"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/shared/service"
)

var _ repository.Content = (*contentRepositoryImpl)(nil)

func NewContentRepository(db *sql.DB) repository.Content {
	return &contentRepositoryImpl{db: db}
}

type contentRepositoryImpl struct {
	db *sql.DB
}

func (c contentRepositoryImpl) Save(ctx context.Context, id uuid.UUID, name string) error {
	if err := service.WithTransaction(ctx, c.db, func(tx *sql.Tx) error {
		q := rdb.New(tx)
		binary, err := id.MarshalBinary()
		if err != nil {
			return err
		}
		req := rdb.CreateContentsParams{
			ID:   binary,
			Name: name,
		}
		if err := q.CreateContents(ctx, req); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
