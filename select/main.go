package main

import (
	"log"
	"time"
)

// 話者３人が会話
func main() {
	log.Println("<main.go>[main]START")

	log.Println("<main.go>[main]トークルームを生成(newTalkroom)します。")
	r := newTalkroom()
	log.Println("<main.go>[main]トークルームを生成(newTalkroom)しました。")
	log.Println("<main.go>[main]トークルームをゴルーチンで起動(runTalkRoom)します。")
	go runTalkRoom(r)
	log.Println("<main.go>[main]トークルームをゴルーチンで起動(runTalkRoom)しました。")
	log.Println("　")

	log.Println("<main.go>[main]話者「Taro」を生成(newTalker)します。")
	t1 := newTalker("Taro", r)
	log.Println("<main.go>[main]話者「Taro」を生成(newTalker)しました。")
	log.Println("<main.go>[main]話者「Taro」をゴルーチンで会話中状態(talking)にします。")
	go t1.talking()
	log.Println("<main.go>[main]話者「Taro」をゴルーチンで会話中状態(talking)にしました。")
	log.Println("　")

	log.Println("<main.go>[main]話者「Jiro」を生成(newTalker)します。")
	t2 := newTalker("Jiro", r)
	log.Println("<main.go>[main]話者「Jiro」を生成(newTalker)しました。")
	log.Println("<main.go>[main]話者「Jiro」をゴルーチンで会話中状態(talking)にします。")
	go t2.talking()
	log.Println("<main.go>[main]話者「Jiro」をゴルーチンで会話中状態(talking)にしました。")
	log.Println("　")

	log.Println("<main.go>[main]話者「Hanako」を生成(newTalker)します。")
	t3 := newTalker("Hanako", r)
	log.Println("<main.go>[main]話者「Hanako」を生成(newTalker)しました。")
	log.Println("<main.go>[main]話者「Hanako」をゴルーチンで会話中状態(talking)にします。")
	go t3.talking()
	log.Println("<main.go>[main]話者「Hanako」をゴルーチンで会話中状態(talking)にしました。")
	log.Println("　")

	log.Println("　")
	log.Println("<main.go>[main]ゴルーチンで適当な３人の会話をスタート")
	log.Println("　")
	// 会話（お試しなので適当に）
	go func() {
		t1.talk <- "Hello!"
		time.Sleep(1 * time.Second)
		t2.talk <- "こんにちは！"
		t3.talk <- "Hi."
		time.Sleep(3 * time.Second)

		t3.talk <- "２個目：さんさん"
		time.Sleep(1 * time.Second)
		t1.talk <- "２個目：いちいち"
		time.Sleep(1 * time.Second)
		t2.talk <- "２個目：にーにー"
	}()

	log.Println("　")
	log.Println("<main.go>[main]会話終了の前に１５秒待機")
	// 適当な時間で会話は強制終了
	time.Sleep(15 * time.Second)
	log.Println("<main.go>[main]会話終了の前に１５秒待機　終わり")
	log.Println("　")

	log.Println("<main.go>[main]話者「Taro」の会話を終了します。")
	t1.bye()
	log.Println("<main.go>[main]話者「Taro」の会話を終了しました。")
	log.Println("　")

	log.Println("<main.go>[main]話者「Jiro」の会話を終了します。")
	t2.bye()
	log.Println("<main.go>[main]話者「Jiro」の会話を終了しました。")
	log.Println("　")

	log.Println("<main.go>[main]話者「Hanako」の会話を終了します。")
	t3.bye()
	log.Println("<main.go>[main]話者「Hanako」の会話を終了しました。")
	log.Println("　")

	time.Sleep(5 * time.Second)
	close(r.joinning)
	close(r.leaving)
	close(r.talk)
	log.Println("<main.go>[main]END")
}
