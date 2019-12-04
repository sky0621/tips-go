package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/profile"
)

func parse(filePath string) *BigJSON {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(f)

	bj := &BigJSON{}
	for {
		t, err := dec.Token()
		if err != nil {
			log.Fatal(err)
		}

		if ele, ok := t.(string); ok {
			if ele == "data" {
				bj.Data = &Data{}
				continue
			}
			if ele == "items" {
				bj.Data.Items = []*Item{}
				continue
			}
		}
		if de, ok := t.(json.Delim); ok {
			if de.String() == "[" {
				break
			}
		}
	}

	for dec.More() {
		var item *Item
		err := dec.Decode(&item)
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

	bj := parse("./big.json")

	fmt.Printf("%f秒\n", time.Now().Sub(st).Seconds())

	for n, item := range bj.Data.Items {
		// Itemsの一番最後の要素だけ試しに確認してみる。
		if n == len(bj.Data.Items)-1 {
			fmt.Printf("[%d] %#v\n", n, item)
		}
	}
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
