package repository

import "github.com/sky0621/tips-go/experiment/architecture_model/internal/content/domain/model"

type Content interface {
	SearchContents(partialName *string) ([]model.Content, error)
}
