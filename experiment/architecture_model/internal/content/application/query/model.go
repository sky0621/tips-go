package query

import "github.com/sky0621/tips-go/experiment/architecture_model/internal/shared/model"

type ProgramReadModel struct {
	ID       model.ID
	Question string
	Answer   string
}
