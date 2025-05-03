package model

import (
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/shared/model"
)

type SaveContentWriteModel struct {
	ID       model.ID
	Name     string
	Programs []ProgramWriteModel
}

type ProgramWriteModel struct {
	ID       model.ID
	Question string
	Answer   string
}
