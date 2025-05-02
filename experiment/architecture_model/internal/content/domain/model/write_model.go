package model

import "github.com/google/uuid"

type SaveContentWriteModel struct {
	ID       uuid.UUID
	Name     string
	Programs []ProgramWriteModel
}

type ProgramWriteModel struct {
	ID       uuid.UUID
	Question string
	Answer   string
}
