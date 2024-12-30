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
	// 使用メモリチェック準備
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	st := time.Now()

	// 125MBのファイルを読み込み
	ba, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatal(err)
	}

	var bj *BigJSON
	// 125MBのデータをパース
	err = json.Unmarshal(ba, &bj)
	if err != nil {
		log.Fatal(err)
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
