package main

import "fmt"

func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func main() {
	prm := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r1, r2 := splitAnySlice(prm)
	fmt.Println(r1)
	fmt.Println(r2)
}
