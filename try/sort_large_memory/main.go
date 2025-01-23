package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// データ数
	const numElements = 100000000

	// ランダムなデータを生成
	data := make([]int, numElements)
	for i := 0; i < numElements; i++ {
		data[i] = rand.Intn(100000000)
	}

	// 並べ替え開始前のタイムスタンプとメモリ使用量を取得
	start := time.Now()

	// データを並べ替える
	sort.Ints(data)

	// 並べ替え完了後のタイムスタンプとメモリ使用量を取得
	elapsed := time.Since(start)

	// 実行時間を表示
	fmt.Printf("Sorted %d elements in %v\n", numElements, elapsed)
}
