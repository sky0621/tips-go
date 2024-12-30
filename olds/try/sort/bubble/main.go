package main

import "fmt"

func main() {
	strs := []string{
		"Go",
		"Java",
		"C++",
		"C",
		"Ruby",
		"Python",
		"PHP",
		"C#",
		"Scala",
		"Kotlin",
		"Rust",
		"Erlang",
	}

	res := bs(strs)

	fmt.Println("")
	fmt.Println(res)
	fmt.Println("")
	for i, str := range res {
		fmt.Printf("[%2d][len:%2d]%s\n", i, len(str), str)
	}
}

func bs(strs []string) []string {
	for i := 0; i < len(strs); i++ {
		for j := i; j < len(strs); j++ {
			if i == j {
				continue
			}
			if len(strs[i]) > len(strs[j]) {
				tmp := strs[i]
				strs[i] = strs[j]
				strs[j] = tmp
			}
		}
	}
	return strs
}
