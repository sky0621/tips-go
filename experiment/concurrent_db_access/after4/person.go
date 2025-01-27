package main

import "errors"

type PersonID interface {
}

type PersonName string

type Person struct {
	ID   PersonID
	Name PersonName
}

func GetPerson(id PersonID) (*Person, DomainError) {
	if id == PersonID(0) {
		return nil, NotFoundError(errors.New("no person found"))
	}
}
