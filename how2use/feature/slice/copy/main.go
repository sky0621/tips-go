package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	var dst []int
	cnt := copy(dst, src)
	fmt.Println("＜スライスをコピー(dstは[]intで宣言しただけではコピーされない)＞", cnt)
	fmt.Printf("[スライス]\t[src]:\t %#v\n", src)
	fmt.Printf("[スライス]\t[dst]:\t %#v\n", dst)
	fmt.Println()

	dst2 := []int{}
	cnt2 := copy(dst2, src)
	fmt.Println("＜スライスをコピー(dst2は[]int{}で空要素初期化しただけではコピーされない)＞", cnt2)
	fmt.Printf("[スライス]\t[src]:\t %#v\n", src)
	fmt.Printf("[スライス]\t[dst2]:\t %#v\n", dst2)
	fmt.Println()

	dst3 := []int{6}
	cnt3 := copy(dst3, src)
	fmt.Println("＜スライスをコピー(dst3は[]int{6}で初期化した場合、サイズ分だけがコピー（入れ替え）される)＞", cnt3)
	fmt.Printf("[スライス]\t[src]:\t %#v\n", src)
	fmt.Printf("[スライス]\t[dst]:\t %#v\n", dst)
	fmt.Println()

	dst4 := make([]int, 0, len(src))
	cnt4 := copy(dst4, src)
	fmt.Println("", cnt4)
	fmt.Printf("[スライス]\t[src]:\t %#v\n", src)
	fmt.Printf("[スライス]\t[dst4]:\t %#v\n", dst4)
	fmt.Println()

}
