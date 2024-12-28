package model

import "fmt"

type UnvalidatedName struct {
	givenName  *GivenName
	familyName *FamilyName
}

func (u *UnvalidatedName) GetValue() string {
	return fmt.Sprintf("%s %s", u.givenName, u.familyName)
}

func NewUnvalidatedName(givenName, familyName string) *UnvalidatedName {
	return &UnvalidatedName{
		givenName:  &GivenName{givenName},
		familyName: &FamilyName{familyName},
	}
}
