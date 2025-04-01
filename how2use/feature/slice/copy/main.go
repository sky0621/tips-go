package main

import "fmt"

// コピー先スライスのサイズはコピー元と合わせないとコピーされない
/*
	<<< bad >>>
	src: [0 1 2]
	dst: []
*/
func bad() {
	fmt.Println("<<< bad >>>")
	src := []int{0, 1, 2}
	var dst []int
	copy(dst, src)
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dst: %v\n", dst)
}

/*
<<< correct >>>
src: [0 1 2]
dst: [0 1 2]
*/
func correct() {
	fmt.Println("<<< correct >>>")
	src := []int{0, 1, 2}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dst: %v\n", dst)
}

/*
<<< correctb >>>
src: [0 1 2]
dst: [0 1 2]
-----
src: [7 1 2]
dst: [0 1 2]
-----
src: [7 1 2]
dst: [0 5 2]
*/
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

/*
<<< correct2 >>>
src: [0 1 2]
dst: [0 1 2]
*/
func correct2() {
	fmt.Println("<<< correct2 >>>")
	src := []int{0, 1, 2}
	dst := make([]int, len(src))
	dst = src[:]
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dst: %v\n", dst)
}

/*
<<< correct2b >>>
src: [0 1 2]
dst: [0 1 2]
-----
src: [7 1 2]
dst: [7 1 2]
-----
src: [7 5 2]
dst: [7 5 2]
*/
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
