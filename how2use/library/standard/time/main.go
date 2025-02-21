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

	log.Println("＜現在日時を（用意された）いろんなフォーマットで表示＞")
	log.Println("[ANSIC]", now.Format(time.ANSIC))
	log.Println("[UnixDate]", now.Format(time.UnixDate))
	log.Println("[RubyDate]", now.Format(time.RubyDate))
	log.Println("[RFC822]", now.Format(time.RFC822))
	log.Println("[RFC822Z]", now.Format(time.RFC822Z))
	log.Println("[RFC850]", now.Format(time.RFC850))
	log.Println("[RFC1123]", now.Format(time.RFC1123))
	log.Println("[RFC1123Z]", now.Format(time.RFC1123Z))
	log.Println("[RFC3339]", now.Format(time.RFC3339))
	log.Println("[RFC3339Nano]", now.Format(time.RFC3339Nano))
	log.Println("[Kitchen]", now.Format(time.Kitchen))
	log.Println("[Stamp]", now.Format(time.Stamp))
	log.Println("[StampMilli]", now.Format(time.StampMilli))
	log.Println("[StampMicro]", now.Format(time.StampMicro))
	log.Println("[StampNano]", now.Format(time.StampNano))
}
