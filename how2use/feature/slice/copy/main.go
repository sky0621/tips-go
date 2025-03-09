package main

import "fmt"

// コピー先スライスのサイズはコピー元と合わせないとコピーされない
func bad() {
	fmt.Println("<<< bad >>>")
	src := []int{0, 1, 2}
	var dst []int
	copy(dst, src)
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dst: %v\n", dst)
}

func correct() {
	fmt.Println("<<< correct >>>")
	src := []int{0, 1, 2}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dst: %v\n", dst)
}
func correctb() {
	fmt.Println("<<< correctb >>>")
	src := []int{0, 1, 2}
	dst := make([]int, len(src))
	/*
	 * copyすると参照先の配列自体もコピーされる
	 */
	copy(dst, src)
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dst: %v\n", dst)
	fmt.Println("-----")
	src[0] = 7
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dst: %v\n", dst)
	fmt.Println("-----")
	dst[1] = 5
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dst: %v\n", dst)
}

func correct2() {
	fmt.Println("<<< correct2 >>>")
	src := []int{0, 1, 2}
	dst := make([]int, len(src))
	dst = src[:]
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dst: %v\n", dst)
}
func correct2b() {
	fmt.Println("<<< correct2b >>>")
	src := []int{0, 1, 2}
	dst := make([]int, len(src))
	/*
	 * スライシングすると参照先の配列自体はコピーされない
	 */
	dst = src[:]
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dst: %v\n", dst)
	fmt.Println("-----")
	src[0] = 7
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dst: %v\n", dst)
	fmt.Println("-----")
	dst[1] = 5
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dst: %v\n", dst)
}

func main() {
	bad()
	fmt.Println()
	correct()
	fmt.Println()
	correctb()
	fmt.Println()
	correct2()
	fmt.Println()
	correct2b()
}
