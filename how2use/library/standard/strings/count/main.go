package main

import (
	"fmt"
	"unicode/utf8"
)

// string はバイト列
// rune は文字
func main() {
	str := "あいうえお"
	fmt.Println(str)
	fmt.Printf("len(str): %d\n", len(str)) // 15

	fmt.Println()

	for i, r := range str {
		// 0: 'あ' (U+3042)
		// 3: 'い' (U+3044)
		// 6: 'う' (U+3046)
		// 9: 'え' (U+3048)
		// 12: 'お' (U+304A)
		fmt.Printf("%d: %q (%[2]U)\n", i, r)
	}
	fmt.Printf("utf8.RuneCountInString(str): %d\n", utf8.RuneCountInString(str)) // 5

	fmt.Println()

	runes := []rune(str)
	fmt.Println(runes)
	for i, r := range runes {
		fmt.Printf("%d: %q (%[2]U)\n", i, r)
	}
	fmt.Printf("len([]rune(str)): %d\n", len(runes)) // 5
}
