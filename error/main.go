package main

import (
	"errors"
	"log"
)

func main() {
	err := doSome()
	if err != nil {
		log.Println(err)
	}

	err = doAny()
	if err != nil {
		log.Println(err)
	}
}

func doSome() error {
	return errors.New("Normal error")
}

func doAny() error {
	return OwnError{errID: 9, errMsg: "Own error"}
}
