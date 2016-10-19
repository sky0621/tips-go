package main

import "log"

func main() {
	defer log.Println("defer 01")
	log.Println("main 01")
	defer log.Println("defer 02")
	log.Println("main 02")
	defer log.Println("defer 03")
	log.Println("main 03")
	defer log.Println("defer 04")
	if err := recover(); err != nil {
		log.Println("recover 01(呼ばれない): ", err)
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover 02 [panic発生前に defer 定義した場合は呼ばれる] : ", err)
		}
	}()

	// ★パニックを挟む★
	doPanic()
	// ★パニックを挟む★

	if err := recover(); err != nil {
		log.Println("recover 03(呼ばれない): ", err)
	}
	defer log.Println("defer 05 : panic起きる関数の後に記述された defer は呼ばれない・・・。")
}

func doPanic() {
	panic("Panic!")
}
