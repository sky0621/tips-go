package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{"value": 123}`
	var m map[string]any
	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}
	/*
	 * 期待: 型：int, 値：123
	 *
	 * 実際: 型：float64, 値：123
	 */
	fmt.Printf("型：%T, 値：%#v\n", m["value"], m["value"])

	// 以下のようにすれば int に変換することもできる
	iValue := int(m["value"].(float64))
	fmt.Printf("型：%T, 値：%#v\n", iValue, iValue)
}
