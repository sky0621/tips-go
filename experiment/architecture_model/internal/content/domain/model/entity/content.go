package entity

import "github.com/google/uuid"

type Content struct {
	ID   uuid.UUID
	Name string
}

func (c Content) IsEmpty() bool {
	return c.ID == uuid.Nil && c.Name == ""
}
