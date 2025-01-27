package main

import (
	"fmt"

	"github.com/sky0621/tips-go/experiment/value_object/person"
)

func main() {
	p, err := person.New(3, "Saburo")
	if err != nil {
		panic(err)
	}
	fmt.Println(p.IDValue())
	fmt.Println(p.NameValue())

	p2, err2 := person.New(0, "Saburo")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(p2.IDValue())
	}

	p3, err3 := person.New(5, "")
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(p3.IDValue())
	}
}
