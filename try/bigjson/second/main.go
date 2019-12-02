package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/profile"
)

func main() {
	st := time.Now()

	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	fmt.Println("Start")

	f, err := os.Open("../big.json")
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(f)
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("<Open> %T: %v\n", t, t)
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("<Open> %T: %v\n", t, t)
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("<Open> %T: %v\n", t, t)
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("<Open> %T: %v\n", t, t)
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("<Open> %T: %v\n", t, t)

	var items []*Item
	for dec.More() {
		var item *Item
		err := dec.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
		//s, _ := dec.Token()
		// do something with s which is the newly read string
		//fmt.Printf("read %q\n", s)
	}
	//t, err = dec.Token()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("<Close> %T: %v\n", t, t)

	fmt.Printf("%f秒\n", time.Now().Sub(st).Seconds())
	fmt.Println("Unmarshaled")

	//for n, item := range items {
	//	if n == 100 {
	//		fmt.Printf("[%d] %#v\n", n, item)
	//	}
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
