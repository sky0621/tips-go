package main

import "fmt"

func main() {
	sss := []SomeStruct{
		{ID: 15, Name: "Alice", Age: 20},
		{ID: 2, Name: "Bob", Age: 22},
		{ID: 33, Name: "Charlie", Age: 23},
	}

	ids := extract(sss, func(s SomeStruct) int64 { return s.ID })
	fmt.Println(ids)

	names := extract(sss, func(s SomeStruct) string { return s.Name })
	fmt.Println(names)

	ages := extract(sss, func(s SomeStruct) int { return s.Age })
	fmt.Println(ages)
}

type SomeStruct struct {
	ID   int64
	Name string
	Age  int
}

// extract は、任意の型 T のスライスから R 型のフィールドを抽出する汎用関数です。
func extract[T any, R any](items []T, selector func(T) R) []R {
	result := make([]R, len(items))
	for i, item := range items {
		result[i] = selector(item)
	}
	return result
}
