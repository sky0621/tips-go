package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Check01 - NG")
	handleData("TEST")
	fmt.Println()

	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Check02 - OK")
	handleData(file)
	fmt.Println()

	fmt.Println("Check03")
	handleData2(345)
	fmt.Println()

	fmt.Println("Check04")
	handleData2("DUMMYYYYY")
	fmt.Println()

	fmt.Println("Check05")
	handleData2(dummyStruct{name: "Who"})
	fmt.Println()
}

type dummyStruct struct {
	name string
}

func handleData(x interface{}) {
	// f := x.(*os.File)
	f, ok := x.(*os.File)
	if ok {
		if f == nil {
			fmt.Println("File is nil!")
		} else {
			fmt.Println("file is " + f.Name())
		}
	} else {
		fmt.Printf("Not file : %#+v\n", x)
	}
}

func handleData2(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Printf("int : %d\n", x)
	case string:
		fmt.Printf("string : %s\n", x)
	default:
		fmt.Println("others ... ")
	}
}
