package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

func main() {
	//smpl()
	smplFull()
}

func smplFull() {
	var tomlConfig tomlConfig
	_, err := toml.DecodeFile("samplefull.toml", &tomlConfig)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%#+v", tomlConfig)
}

type tomlConfig struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

// ===================================================================================================

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
