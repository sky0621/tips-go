package person

import (
	"fmt"
)

type ID struct {
	v ValueObject[int64]
}

func NewID(v int64) (ID, error) {
	id, err := NewValueObject[int64](v, isID)
	if err != nil {
		return ID{}, fmt.Errorf("idの生成に失敗しました。[%d] %w", v, err)
	}
	return ID{v: id}, nil
}

type Name struct {
	v ValueObject[string]
}

func NewName(v string) (Name, error) {
	name, err := NewValueObject[string](v, isName, isPersonName)
	if err != nil {
		return Name{}, fmt.Errorf("nameの生成に失敗しました。[%s] %w", v, err)
	}
	return Name{v: name}, nil
}

func isPersonName(v string) bool {
	return v != "-"
}

type Person struct {
	id   ID
	name Name
}

func (p Person) ID() ID {
	return p.id
}
func (p Person) IDValue() int64 {
	return p.id.v.Value()
}

func (p Person) Name() Name {
	return p.name
}
func (p Person) NameValue() string {
	return p.name.v.Value()
}

func New(id int64, name string) (Person, error) {
	pID, err := NewID(id)
	if err != nil {
		return Person{}, fmt.Errorf("personの生成に失敗しました。%w", err)
	}
	pName, err := NewName(name)
	if err != nil {
		return Person{}, fmt.Errorf("personの生成に失敗しました。%w", err)
	}
	return Person{
		id:   pID,
		name: pName,
	}, nil
}
