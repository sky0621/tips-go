package main

import "fmt"

func main() {
	/*
	 * 数字配列
	 */
	// 配列は要素数固定
	intArr := [2]int{}
	fmt.Printf("intArr: %v, len: %d\n", intArr, len(intArr))
	fmt.Println()

	intArr[0] = 100
	fmt.Printf("intArr: %v, len: %d\n", intArr, len(intArr))

	// 無効な 配列 インデックス '2' (要素数 2 の配列の範囲外です)
	//intArr[2] = 100

	fmt.Println()
	fmt.Println()

	/*
	 * 文字列配列
	 */
	strArr := [2]string{}
	fmt.Printf("strArr: %v, len: %d\n", strArr, len(strArr))
	fmt.Println()

	// 'strArr' (型 [2]string) を型 []Type として使用できません
	//strArr = append(strArr, "A")
}
