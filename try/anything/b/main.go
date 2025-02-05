package main

import (
	"fmt"
	"time"
)

func main() {
	u := User{
		ID:   1,
		Name: "User1",
		Audit: Audit{
			CreatedBy: 12,
			CreatedAt: time.Now(),
		},
	}
	fmt.Println(u.IsUpdated())
	fmt.Println(u.IsValid())

	u2 := User{
		ID:   0,
		Name: "User1",
		Audit: Audit{
			CreatedBy: 12,
			CreatedAt: time.Now(),
			UpdatedBy: 45,
			UpdatedAt: time.Now(),
		},
	}
	fmt.Println(u2.IsUpdated())
	fmt.Println(u2.IsValid())
}

type Audit struct {
	CreatedBy int
	CreatedAt time.Time
	UpdatedBy int
	UpdatedAt time.Time
}

func (a *Audit) IsUpdated() bool {
	return !a.UpdatedAt.IsZero()
}

type User struct {
	ID   int
	Name string
	Audit
}

func (u *User) IsValid() bool {
	return u.ID > 0
}
