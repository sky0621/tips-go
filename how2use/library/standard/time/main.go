package main

import (
	"log"
	"time"
)

func main() {
	now := time.Now()
	// 現在日時： 2025-04-17 21:45:44.553937 +0900 JST m=+0.000041710
	log.Println("現在日時：", now)

	unow := now.Unix()
	// UNIXタイムに変換： 1744893944
	log.Println("UNIXタイムに変換：", unow)
	// 再びTimeに変換(ナノ無し)： 2025-04-17 21:45:44 +0900 JST
	log.Println("再びTimeに変換(ナノ無し)：", time.Unix(unow, 0))

	utcnow := now.UTC()
	// UTCに変換： 2025-04-17 12:45:44.553937 +0000 UTC
	log.Println("UTCに変換：", utcnow)

	log.Println("＜現在日時を（用意された）いろんなフォーマットで表示＞")
	// [ANSIC]       Thu Apr 17 21:45:44 2025
	log.Println("[ANSIC]", now.Format(time.ANSIC))
	// [UnixDate]    Thu Apr 17 21:45:44 JST 2025
	log.Println("[UnixDate]", now.Format(time.UnixDate))
	// [RubyDate]    Thu Apr 17 21:45:44 +0900 2025
	log.Println("[RubyDate]", now.Format(time.RubyDate))
	// [RFC822]      17 Apr 25 21:45 JST
	log.Println("[RFC822]", now.Format(time.RFC822))
	// [RFC822Z]     17 Apr 25 21:45 +0900
	log.Println("[RFC822Z]", now.Format(time.RFC822Z))
	// [RFC850]      Thursday, 17-Apr-25 21:45:44 JST
	log.Println("[RFC850]", now.Format(time.RFC850))
	// [RFC1123]     Thu, 17 Apr 2025 21:45:44 JST
	log.Println("[RFC1123]", now.Format(time.RFC1123))
	// [RFC1123Z]    Thu, 17 Apr 2025 21:45:44 +0900
	log.Println("[RFC1123Z]", now.Format(time.RFC1123Z))
	// [RFC3339]     2025-04-17T21:45:44+09:00
	log.Println("[RFC3339]", now.Format(time.RFC3339))
	// [RFC3339Nano] 2025-04-17T21:45:44.553937+09:00
	log.Println("[RFC3339Nano]", now.Format(time.RFC3339Nano))
	// [Kitchen]     9:45PM
	log.Println("[Kitchen]", now.Format(time.Kitchen))
	// [Stamp]       Apr 17 21:45:44
	log.Println("[Stamp]", now.Format(time.Stamp))
	// [StampMilli]  Apr 17 21:45:44.553
	log.Println("[StampMilli]", now.Format(time.StampMilli))
	// [StampMicro]  Apr 17 21:45:44.553937
	log.Println("[StampMicro]", now.Format(time.StampMicro))
	// [StampNano]   Apr 17 21:45:44.553937000
	log.Println("[StampNano]", now.Format(time.StampNano))
}
