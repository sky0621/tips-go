package person

type ID struct {
	v ValueObject[int64]
}

func NewID(v int64) (ID, error) {
	id, err := NewValueObject[int64](v, isID)
	if err != nil {
		return ID{}, err
	}
	return ID{v: id}, nil
}

type Name struct {
	v ValueObject[string]
}

func NewName(v string) (Name, error) {
	name, err := NewValueObject[string](v, isName, isPersonName)
	if err != nil {
		return Name{}, err
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
		return Person{}, err
	}
	pName, err := NewName(name)
	if err != nil {
		return Person{}, err
	}
	return Person{
		id:   pID,
		name: pName,
	}, nil
}
