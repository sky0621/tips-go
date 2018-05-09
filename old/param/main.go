package main

import "log"

func main() {
	str := "Before"
	ret1 := test(str)
	log.Println(ret1, str)
	log.Println("===================================")
	ret2 := test2(&str)
	log.Println(*ret2, str)
}

// コピーされた str が渡される
func test(str string) string {
	str = "[" + str + "]"
	return str
}

// str の参照が渡される
func test2(str *string) *string {
	*str = "[" + *str + "]"
	return str
}
