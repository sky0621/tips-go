package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

func main() {
	smpl()
}

func smpl() {
	var da []byte
	var smpl Sample
	var err error
	da, err = ioutil.ReadFile("sample.toml")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(da))

	_, err = toml.DecodeFile("sample.toml", &smpl)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%#v\n", smpl)
	log.Println("============================================")
	d := smpl.Owner.Dob
	log.Printf("%04d/%02d/%02d %02d:%02d:%02d\n", d.Year(), d.Month(), d.Day(), d.Hour(), d.Minute(), d.Second())
}

// Sample ...
type Sample struct {
	Title string
	Owner Owner
}

// Owner ...
type Owner struct {
	Name string
	Dob  time.Time
}
