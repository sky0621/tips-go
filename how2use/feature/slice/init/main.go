package main

import "fmt"

// スライスは要素の集合体。配列に似ているが厳密には違う。
// スライスは参照型。（マップやチャネルも参照型）

func main() {
	fmt.Println("＜初期化の違い＞")
	var iSlice []int
	fmt.Println("var iSlice []int")
	fmt.Printf("  iSlice: %v, len: %d\n", iSlice, len(iSlice))

	iSlice2 := []int{}
	fmt.Println("iSlice2 []int")
	fmt.Printf("  iSlice2: %v, len: %d\n", iSlice2, len(iSlice2))

	iSlice2b := []int(nil)
	fmt.Println("iSlice2b []int(nil)")
	fmt.Printf("  iSlice2b: %v, len: %d\n", iSlice2b, len(iSlice2b))

	iSlice3 := make([]int, 0)
	fmt.Println("iSlice3 := make([]int, 0)")
	fmt.Printf("  iSlice3: %v, len: %d\n", iSlice3, len(iSlice3))

	fmt.Println()

	fmt.Println("＜サイズを指定して初期化＞")
	iSlice4 := make([]int, 2, 3)
	fmt.Println("iSlice4 := make([]int, 2, 3)")
	fmt.Printf("  iSlice4: %v, len: %d\n", iSlice4, len(iSlice4))
	iSlice4 = append(iSlice4, 1)
	fmt.Println("iSlice4 = append(iSlice4, 1)")
	fmt.Printf("  iSlice4: %v, len: %d\n", iSlice4, len(iSlice4))
	iSlice4 = append(iSlice4, 2)
	fmt.Println("iSlice4 = append(iSlice4, 2)")
	fmt.Printf("  iSlice4: %v, len: %d\n", iSlice4, len(iSlice4))
}
