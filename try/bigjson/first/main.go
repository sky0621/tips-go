package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/pkg/profile"
)

func main() {
	st := time.Now()

	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	fmt.Println("Start")

	ba, err := ioutil.ReadFile("../big.json")
	if err != nil {
		log.Fatal(err)
	}
	var bj *BigJSON
	err = json.Unmarshal(ba, &bj)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%f秒\n", time.Now().Sub(st).Seconds())
	fmt.Println("Unmarshaled")

	//for n, item := range bj.Data.Items {
	//	fmt.Printf("[%d] %#v\n", n, item)
	//}

	//fmt.Printf("%f秒\n", time.Now().Sub(st).Seconds())
	fmt.Println("End")
}

type BigJSON struct {
	Data *Data
}

type Data struct {
	Items []*Item
}

type Item struct {
	ID      string
	Name    string
	Price   int
	Display bool
}
