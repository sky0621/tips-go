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

	fmt.Println()

	fmt.Println("＜appendして別のスライスを生成＞")
	iSliceOrg := make([]int, 0, 3)
	iSliceOrg = append(iSliceOrg, 1)
	fmt.Println("iSliceOrg := make([]int, 0, 3)")
	fmt.Println("iSliceOrg = append(iSliceOrg, 1)")
	fmt.Printf("  iSliceOrg: %v, len: %d\n", iSliceOrg, len(iSliceOrg))

	fmt.Println()

	// iSliceOrgと同じ配列を指す新しいスライスを生成
	//
	// [1][2]
	//
	// iSliceOrgは上記の配列の０番目のみを示すスライス
	// isSliceNewは上記の配列のすべての範囲を示すスライス
	iSliceNew := append(iSliceOrg, 2)
	fmt.Println("iSliceNew := append(iSliceOrg, 2)")
	fmt.Printf("  iSliceOrg: %v, len: %d\n", iSliceOrg, len(iSliceOrg))
	fmt.Printf("  iSliceNew: %v, len: %d\n", iSliceNew, len(iSliceNew))

	fmt.Println()

	// iSliceOrgとiSliceNewは現時点では同じ配列を示している
	// そのため、どちらかの要素を書き換えた結果は両方のスライスに反映される
	iSliceOrg[0] = 11
	fmt.Println("iSliceOrg[0] = 11")
	fmt.Printf("  iSliceOrg: %v, len: %d\n", iSliceOrg, len(iSliceOrg))
	fmt.Printf("  iSliceNew: %v, len: %d\n", iSliceNew, len(iSliceNew))

	// iSliceNewに要素を追加
	// iSliceOrgは同じ配列の０番目しか見ていないので上記操作の結果はiSliceNewにだけ反映される
	//
	// iSliceOrg: [11], len: 1
	// iSliceNew: [11 2 3], len: 3
	iSliceNew = append(iSliceNew, 3)
	fmt.Println("iSliceNew = append(iSliceNew, 3)")
	fmt.Printf("  iSliceOrg: %v, len: %d\n", iSliceOrg, len(iSliceOrg))
	fmt.Printf("  iSliceNew: %v, len: %d\n", iSliceNew, len(iSliceNew))

	// iSliceOrgに要素を追加
	//
	// iSliceOrgとiSliceNewとで以下の同じ配列を（範囲違いで）参照していて、iSliceOrgは [11] だけを参照しているため、
	// [11][2][3]
	// iSliceOrgに要素を追加すると、結果的に [2] を書き換えることになる（つまり、iSliceNewにも影響が出る）
	//
	// [11][22][3]
	iSliceOrg = append(iSliceOrg, 22)
	fmt.Println("iSliceOrg = append(iSliceOrg, 22)")
	fmt.Printf("  iSliceOrg: %v, len: %d\n", iSliceOrg, len(iSliceOrg))
	fmt.Printf("  iSliceNew: %v, len: %d\n", iSliceNew, len(iSliceNew))

	fmt.Println()

	// iSliceNewにさらに要素を加える
	//
	// これまでに操作していた配列は最初にキャパシティ「３」で作成していたもののため、このサイズに収まっている範囲内の操作は
	// 同じ配列を参照している iSliceOrg と iSliceNew の両方に（参照範囲が同じなら）影響が出る
	//
	// ただ、現在のサイズが３である状態でさらに要素を追加すると、追加するために別の配列を用意する必要が生じ、
	// 結果的に iSliceNew は iSliceOrg とは別の新しい配列を参照することになる
	// これ以降は、参照範囲が同じ場所を修正したとしても参照先の配列が違うため、他方のスライスに影響が出なくなる
	iSliceNew = append(iSliceNew, 4)
	fmt.Println("iSliceNew = append(iSliceNew, 4)")
	fmt.Printf("  iSliceOrg: %v, len: %d\n", iSliceOrg, len(iSliceOrg))
	fmt.Printf("  iSliceNew: %v, len: %d\n", iSliceNew, len(iSliceNew))

	iSliceOrg[0] = 111
	fmt.Println("iSliceOrg[0] = 111")
	fmt.Printf("  iSliceOrg: %v, len: %d\n", iSliceOrg, len(iSliceOrg))
	fmt.Printf("  iSliceNew: %v, len: %d\n", iSliceNew, len(iSliceNew))

	iSliceNew[1] = 222
	fmt.Println("iSliceNew[1] = 222")
	fmt.Printf("  iSliceOrg: %v, len: %d\n", iSliceOrg, len(iSliceOrg))
	fmt.Printf("  iSliceNew: %v, len: %d\n", iSliceNew, len(iSliceNew))
}
