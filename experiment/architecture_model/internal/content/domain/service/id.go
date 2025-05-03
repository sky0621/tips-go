package service

import (
	"github.com/google/uuid"
)

func MustCreateNewID() uuid.UUID {
	u, e := CreateNewID()
	if e != nil {
		panic(e)
	}
	return u
}

func CreateNewID() (uuid.UUID, error) {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return [16]byte{}, err
	}
	return uuidV7, nil
}

func ParseID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func CreateID(id []byte) (uuid.UUID, error) {
	uuidID, err := uuid.FromBytes(id)
	if err != nil {
		return uuid.UUID{}, err
	}
	return uuidID, nil
}
