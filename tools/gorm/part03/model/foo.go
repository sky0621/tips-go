package model

import "github.com/jinzhu/gorm"

type Foo struct {
	gorm.Model
}

func (Foo) TableName() string {
	return "foo_alias"
}

type Baa struct {
	gorm.Model
	Kind int
}

func (b Baa) TableName() string {
	if b.Kind == 1 {
		return "admin_baa"
	} else {
		return "user_baa"
	}
}
