package person

type UnvalidatedID int64

type UnvalidatedName string

type UnvalidatedPerson struct {
	ID   UnvalidatedID
	Name UnvalidatedName
}

func (p UnvalidatedPerson) IDValue() int64    { return int64(p.ID) }
func (p UnvalidatedPerson) NameValue() string { return string(p.Name) }
