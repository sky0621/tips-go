package target

// PGType ...
type PGType int8

// ProgrammingLanguage ...
type ProgrammingLanguage interface {
	GetName() string
	GetPGType() PGType
	Create(name string, pgType PGType)
}
