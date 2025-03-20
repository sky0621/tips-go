package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
 * time.Time は、MarshalJSON でシリアライズすると、
 * モノトニッククロックの情報が失われる。
 */
func main() {
	t1 := time.Now()
	d1, err := json.Marshal(t1)
	if err != nil {
		panic(err)
	}

	var t2 time.Time
	if err := json.Unmarshal(d1, &t2); err != nil {
		panic(err)
	}

	// 2025-03-20 23:36:02.829895 +0900 JST m=+0.000135668
	fmt.Println(t1)
	// 2025-03-20 23:36:22.641636 +0900 JST
	fmt.Println(t2)

	fmt.Println("t1 == t2 ? ...", t1 == t2)
	fmt.Println("t1.Equal(t2) ? ...", t1.Equal(t2))
}
