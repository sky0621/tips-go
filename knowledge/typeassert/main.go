package main

import (
	"log"
	"os"
)

func main() {
	handleData("TEST")
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)
	handleData(file)

	log.Println("============================")

	handleData2(345)
	handleData2("DUMMYYYYY")
	handleData2(dummyStruct{name: "Who"})
}

type dummyStruct struct {
	name string
}

func handleData(x interface{}) {
	// f := x.(*os.File)
	f, ok := x.(*os.File)
	if ok {
		log.Println("file is " + f.Name())
	} else {
		log.Printf("Not file : %#+v\n", x)
	}
}

func handleData2(x interface{}) {
	switch x.(type) {
	case int:
		log.Printf("int : %d\n", x)
	case string:
		log.Printf("string : %s\n", x)
	default:
		log.Println("others ... ")
	}
}
