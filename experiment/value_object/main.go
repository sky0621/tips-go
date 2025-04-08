package main

import (
	"fmt"
	"github.com/sky0621/tips-go/experiment/value_object/person"
)

func main() {
	samplePerson()
}

func samplePerson() {
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

	fmt.Println()

	uvp1 := person.UnvalidatedPerson{
		ID:   99,
		Name: "Siro",
	}
	fmt.Println(uvp1)

	vp1, err4 := person.New(uvp1.IDValue(), uvp1.NameValue())
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(vp1.IDValue(), vp1.NameValue())
	}

	uvp2 := person.UnvalidatedPerson{
		ID:   -1,
		Name: "",
	}
	fmt.Println(uvp2)

	vp2, err5 := person.New(uvp2.IDValue(), uvp2.NameValue())
	if err5 != nil {
		fmt.Println(err5)
	} else {
		fmt.Println(vp2.IDValue(), vp2.NameValue())
	}
}
