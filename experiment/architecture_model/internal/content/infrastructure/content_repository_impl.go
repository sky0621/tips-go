package infrastructure

import (
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/domain/model"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/domain/repository"
)

var _ repository.Content = (*contentRepositoryImpl)(nil)

type contentRepositoryImpl struct {
}

func (c contentRepositoryImpl) SearchContents(partialName *string) ([]model.Content, error) {
	//TODO implement me
	panic("implement me")
}
