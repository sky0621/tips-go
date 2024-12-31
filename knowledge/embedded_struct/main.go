package main

import (
	"fmt"
	"time"
)

func main() {
	a := Admin{
		Audit: Audit{
			createdBy: 23, createdAt: time.Now(),
			updatedBy: 23, updatedAt: time.Now(),
		},
		name: "Takagi",
	}

	o := Operator{
		Audit: Audit{
			createdBy: 51, createdAt: time.Now(),
			updatedBy: 51, updatedAt: time.Now(),
		},
		name: "Sato",
		roles: []Role{
			{id: 3, name: "AreaManager"},
		},
	}

	g := Guest{
		Audit: Audit{
			createdBy: 109, createdAt: time.Now(),
			updatedBy: 109, updatedAt: time.Now(),
		},
		id: 617,
	}

	fmt.Println(a)
	fmt.Println(o)
	fmt.Println(g)
}

type Audit struct {
	createdBy int
	createdAt time.Time
	updatedBy int
	updatedAt time.Time
}

func (a Audit) String() string {
	return fmt.Sprintf("Audit{createdBy:%d, createdAt:%v, updatedBy:%d, updatedAt:%v",
		a.createdBy, a.createdAt, a.updatedBy, a.updatedAt)
}

type Admin struct {
	Audit
	name string
}

func (a Admin) String() string {
	return fmt.Sprintf("Admin{name:%s, audit:%+v}", a.name, a.Audit)
}

type Operator struct {
	Audit
	name  string
	roles []Role
}

func (o Operator) String() string {
	return fmt.Sprintf("Operator{name:%s, roles:%+v, audit:%+v}", o.name, o.roles, o.Audit)
}

type Guest struct {
	Audit
	id int
}

func (g Guest) String() string {
	return fmt.Sprintf("Guest{id:%d, audit:%+v}", g.id, g.Audit)
}

type Role struct {
	id   int
	name string
}

func (r Role) String() string {
	return fmt.Sprintf("Role{id:%d, name:%s}", r.id, r.name)
}
