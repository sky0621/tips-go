package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now()
	log.Println("現在日時：", now)
	unow := now.Unix()
	log.Println("UNIXタイムに変換：", unow)
	log.Println("再びTimeに変換(ナノ無し)：", time.Unix(unow, 0))
	utcnow := now.UTC()
	log.Println("UTCに変換：", utcnow)

	log.Println("PM:", now.Format("3:04PM"))

	str := "Thu May 24 22:56:30 JST 2001"
	layout := "Mon Jan 2 15:04:05 MST 2006"
	t, _ := time.Parse(layout, str)
	fmt.Println(t) // => "2001-05-24 22:56:30 +0900 JST"

	str = "2003/04/18"
	layout = "2006/01/02"
	t, _ = time.Parse(layout, str)
	fmt.Println(t) // => "2003-04-18 00:00:00 +0000 UTC"
}
