package main

import "log"

func main() {
	defer log.Println("[main] 最初に定義した defer")

	log.Println("[main] 01")

	defer log.Println("[main] ２番目に定義した defer")

	log.Println("[main] 02")

	defer log.Println("[main] ３番目に定義した defer")

	log.Println("[main] 03")

	if err := recover(); err != nil {
		log.Println("[main] 呼ばれない recover: ", err)
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("[main] panic発生前に defer 定義したので呼ばれる recover: ", err)
		}
	}()

	// ★パニックを挟む★
	doPanic()
	// ★パニックを挟む★

	defer func() {
		if err := recover(); err != nil {
			log.Println("[main] panic発生後に defer 定義したので呼ばれない recover: ", err)
		}
	}()

	if err := recover(); err != nil {
		log.Println("[main] panic発生後に定義したので呼ばれない recover: ", err)
	}

	defer log.Println("[main] panic発生後に記述された defer は呼ばれない・・・。")
}

func doPanic() {
	panic("Panic!")
}
