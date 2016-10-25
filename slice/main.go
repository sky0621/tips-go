package main

import "fmt"

// スライスは要素の集合体。配列に似ているが厳密には違う。
// スライスは参照型。（マップやチャネルも参照型）

func main() {
	// 文字列を要素とするスライス。参照型なのでこの時点では nil
	var s []string

	// 配列の宣言と初期化
	a := [5]string{"Go", "Java", "Scala", "Ruby", "Python"}

	// 配列のすべてをスライスへ
	s = a[:]
	fmt.Println("＜５つの文字列配列＞")
	fmt.Printf("[スライス]\t[s]:\t %#v\n", s)
	fmt.Printf("[配列]\t\t[a]:\t%#v\n", a)
	fmt.Println()

	s[2] = "PHP"
	fmt.Println("＜スライスの３番目をPHPに変更＞")
	fmt.Printf("[スライス]\t[s]:\t %#v\n", s)
	fmt.Printf("[配列]\t\t[a]:\t%#v\n", a)
	fmt.Println()

	s2 := append(s, "Elixir", "C++")
	fmt.Println("＜スライスに要素「Elixir」「C++」を追加(★元のスライス、配列には影響なし★) -> s2 ＞")
	fmt.Printf("[スライス]\t[s2]:\t %#v\n", s2)
	fmt.Printf("[スライス]\t[s]:\t %#v\n", s)
	fmt.Printf("[配列]\t\t[a]:\t%#v\n", a)
	fmt.Println()

	src := []int{1, 2, 3, 4, 5}
	var dst []int
	cnt := copy(dst, src)
	fmt.Println("＜スライスをコピー(dstは[]intで宣言しただけではコピーされない)＞", cnt)
	fmt.Printf("[スライス]\t[src]:\t %#v\n", src)
	fmt.Printf("[スライス]\t[dst]:\t %#v\n", dst)
	fmt.Println()

	dst = []int{}
	cnt = copy(dst, src)
	fmt.Println("＜スライスをコピー(dstは[]int{}で空要素初期化しただけではコピーされない)＞", cnt)
	fmt.Printf("[スライス]\t[src]:\t %#v\n", src)
	fmt.Printf("[スライス]\t[dst]:\t %#v\n", dst)
	fmt.Println()

	dst = []int{6}
	cnt = copy(dst, src)
	fmt.Println("＜スライスをコピー(dstは[]int{6}で初期化した場合、サイズ分だけがコピー（入れ替え）される)＞", cnt)
	fmt.Printf("[スライス]\t[src]:\t %#v\n", src)
	fmt.Printf("[スライス]\t[dst]:\t %#v\n", dst)
	fmt.Println()

	// いったん配列として初期化
	dst2 := [...]int{}
	cnt = copy(dst2[:], src)
	fmt.Println("＜スライスをコピー(dst2は、いったん[...]int{}で配列として初期化した場合、結局サイズ０なのでコピーされない)＞", cnt)
	fmt.Printf("[スライス]\t[src]:\t %#v\n", src)
	fmt.Printf("[スライス]\t[dst2]:\t %#v\n", dst2)
	fmt.Println()

	isl := make([]int, 10, 10)
	fmt.Println("＜スライスをmake([]int, 10, 10)で生成＞")
	fmt.Println(isl)
	fmt.Println()

	bt := []string{"A", "B", "O", "AB"}
	fmt.Println("＜スライスを可変長引数持ち関数に渡す＞")
	test(bt...)

}

// 可変長引数を持つ関数
func test(s ...string) {
	fmt.Println(len(s), s)
	fmt.Println()
}
