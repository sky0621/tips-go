package main

import (
	"fmt"
	"strings"
)

func main() {
	fn1()
	fmt.Println()
	fn2()
}

func fn1() {
	courses := []string{"A", "B", "C", "D", "E"}
	for i, c := range courses {
		fmt.Printf("%d: %s\n", i, c)
		if i%2 == 0 {
			c = strings.ToLower(c)
		}
	}
	fmt.Println("-----")
	fmt.Println(courses)
}

func fn2() {
	courses := []string{"A", "B", "C", "D", "E"}
	for i, c := range courses {
		fmt.Printf("%d: %s\n", i, c)
		if i%2 == 0 {
			courses[i] = strings.ToLower(c)
		}
	}
	fmt.Println("-----")
	fmt.Println(courses)
}
