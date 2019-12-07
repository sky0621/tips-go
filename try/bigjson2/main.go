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
	// 使用メモリチェック準備
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	st := time.Now()

	f, err := os.Open("./data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	dec := json.NewDecoder(f)

	bj := &BigJSON{}
	// 以下のJSON構造をトークン毎（「 { 」、「 "data" 」、「 { 」、「 "items" 」、「 [ 」、・・・）に辿る。
	// ★「 [ 」以降は Item 構造体にパースするObject構造が複数続くので、後続の dec.Decode() に任せる。
	// {
	//   "data": {
	//     "items": [
	//       {
	//        "id": "id-0",
	//        "name": "あいてむ0",
	//        "price": 0,
	//        "display": true
	//      },
	//        〜〜省略〜〜
	//      {
	//        "id": "id-999999",
	//        "name": "あいてむ999999",
	//        "price": 999999,
	//        "display": true
	//      }
	//    ]
	//  }
	//}
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
				// ここまで辿れば、後は後続の dec.Decode() に任せるため break
				break
			}
		}
	}

	// ここに来た時点で、Item 構造体にパース可能なObject構造が複数続くので、ある分だけひたすらパース
	for dec.More() {
		var item *Item
		err := dec.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
		bj.Data.Items = append(bj.Data.Items, item)
	}

	fmt.Printf("%f秒\n", time.Now().Sub(st).Seconds())
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
