package entity

import "github.com/google/uuid"

type Program struct {
	ID       uuid.UUID
	Question string
	Answer   string
}
