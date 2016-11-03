package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

func main() {
	smpl01()
	smpl02()
}

// 事例１

func smpl01() {
	var smpl01 Sample01
	var err error
	_, err = toml.DecodeFile("sample01.toml", &smpl01)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%#v", smpl01)
}

// Sample01 ...
type Sample01 struct {
	Title string
}

// 事例２

func smpl02() {
	var smpl02 Sample02
	var err error
	_, err = toml.DecodeFile("sample02.toml", &smpl02)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%#v", smpl02)
}

// Sample02 ...
type Sample02 struct {
	Title string
	Owner Owner
}

// Owner ...
type Owner struct {
	Name string
}
