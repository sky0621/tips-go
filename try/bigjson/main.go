package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/pkg/profile"
)

func parse(filePath string) *BigJSON {
	ba, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var bj *BigJSON
	err = json.Unmarshal(ba, &bj)
	if err != nil {
		log.Fatal(err)
	}

	return bj
}

func parse2(filePath string) *BigJSON {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(f)

	bj := &BigJSON{}
	var isItems bool

	for dec.More() {
		t, err := dec.Token()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[%T] %v\n", t, t)

		if ele, ok := t.(string); ok {
			if ele == "data" {
				bj.Data = &Data{}
				continue
			}
			if ele == "items" {
				bj.Data.Items = []*Item{}
				isItems = true
				continue
			}
		}

		if !isItems {
			continue
		}

		if de, ok := t.(json.Delim); ok {
			if de.String() == "[" {
				continue
			}
		}

		var item *Item
		err = dec.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
		bj.Data.Items = append(bj.Data.Items, item)
	}
	return bj
}

func main() {
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	fmt.Println("Start")

	st := time.Now()

	bj := parse2("/home/sky0621/work/src/go111/src/github.com/sky0621/tips-go/try/bigjson/big.json")

	fmt.Printf("%f秒\n", time.Now().Sub(st).Seconds())

	for n, item := range bj.Data.Items {
		if n == len(bj.Data.Items) {
			fmt.Printf("[%d] %#v\n", n, item)
		}
	}
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
