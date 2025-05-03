package service

import "github.com/google/uuid"

func MustCreateID() uuid.UUID {
	u, e := CreateID()
	if e != nil {
		panic(e)
	}
	return u
}

func CreateID() (uuid.UUID, error) {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return [16]byte{}, err
	}
	return uuidV7, nil
}

func ParseID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}
