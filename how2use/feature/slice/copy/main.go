package main

import "fmt"

// コピー先スライスのサイズはコピー元と合わせないとコピーされない
func bad() {
	src := []int{0, 1, 2}
	var dst []int
	copy(dst, src)
	fmt.Println(dst)
}

func correct() {
	src := []int{0, 1, 2}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println(dst)
}

func main() {
	bad()
	correct()
}
