package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

func main() {
	var smpl01 Sample01
	var err error
	_, err = toml.DecodeFile("sample01.toml", &smpl01)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%#v", smpl01)
}

type Sample01 struct {
	Title string
}
