package main

import (
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
}
