package main

import (
	"fmt"
	"restudy/user"
)

func main() {
	u1 := user.NewUser(5, "Jiro")
	fmt.Println(u1)

}

var _ Item = (*item)(nil)

type Item interface {
	GetItem() string
	AddItem(string)
}

type item struct {
}

func (i item) AddItem(s string) {
	//TODO implement me
	panic("implement me")
}

func (i item) GetItem() string {
	//TODO implement me
	panic("implement me")
}
