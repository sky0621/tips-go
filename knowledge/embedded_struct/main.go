package main

import (
	"fmt"
	"time"
)

func main() {
	crdAt := time.Now()
	updAt := crdAt.Add(time.Hour)

	audit := Audit{
		createdBy: 23, createdAt: crdAt,
		updatedBy: 23, updatedAt: updAt,
	}

	admin := Admin{
		Audit: audit,
		name:  "Takagi",
	}

	operator := Operator{
		Audit: audit,
		name:  "Sato",
		roles: []Role{
			{id: 3, name: "AreaManager"},
		},
	}

	guest := Guest{
		id: 617,
		Audit: Audit{
			createdBy: 109, createdAt: crdAt.Add(10 * time.Minute),
			updatedBy: 131, updatedAt: updAt.Add(10 * time.Minute),
		},
	}

	/*
	   Admin{
	    name:Takagi,
	    audit:Audit{
	     createdBy:23,
	     createdAt:2025-04-01 23:59:01.975034 +0900 JST m=+0.000050459,
	     updatedBy:23,
	     updatedAt:2025-04-02 00:59:01.975034 +0900 JST m=+3600.000050459
	    }
	   }
	*/
	fmt.Println(admin)
	/*
	  Operator{
	   name:Sato,
	   roles:[Role{id:3, name:AreaManager}],
	   audit:Audit{
	    createdBy:23,
	    createdAt:2025-04-01 23:59:01.975034 +0900 JST m=+0.000050459,
	    updatedBy:23,
	    updatedAt:2025-04-02 00:59:01.975034 +0900 JST m=+3600.000050459
	   }
	  }
	*/
	fmt.Println(operator)
	/*
	  Guest{
	   id:617,
	   audit:Audit{
	    createdBy:109,
	    createdAt:2025-04-02 00:09:01.975034 +0900 JST m=+600.000050459,
	    updatedBy:131,
	    updatedAt:2025-04-02 01:09:01.975034 +0900 JST m=+4200.000050459
	   }
	  }
	*/
	fmt.Println(guest)
	fmt.Println(operator.isAudit()) // true
	fmt.Println(guest.isAudit())    // false
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

func (a Audit) isAudit() bool {
	return true
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

func (g Guest) isAudit() bool {
	return false
}

type Role struct {
	id   int
	name string
}

func (r Role) String() string {
	return fmt.Sprintf("Role{id:%d, name:%s}", r.id, r.name)
}
