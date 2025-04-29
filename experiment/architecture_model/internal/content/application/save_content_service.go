package application

import (
	"context"
	"github.com/google/uuid"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/domain/repository"
)

func NewSaveContentService(r repository.Content) SaveContentService {
	return SaveContentService{r: r}
}

type SaveContentService struct {
	r repository.Content
}

func (s SaveContentService) Save(ctx context.Context, id uuid.UUID, name string) error {
	if err := s.r.Save(ctx, id, name); err != nil {
		return err
	}
	return nil
}
