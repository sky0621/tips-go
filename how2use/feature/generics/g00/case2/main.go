package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	strNums := MapIntsToStrings(nums, func(n int) string {
		return fmt.Sprintf("Number: %d", n)
	})
	fmt.Println("Mapped strings:", strNums)

	strs := []string{"aa", "b", "ccc", "dddddd"}
	strInts := MapStringsToInts(strs, func(s string) int {
		return len(s)
	})
	fmt.Println("Mapped ints:", strInts)

	strNums2 := Map(nums, func(n int) string {
		return fmt.Sprintf("Number: %d", n*2)
	})
	fmt.Println("Mapped strings:", strNums2)

	strsInts2 := Map(strs, func(s string) int {
		return len(s) * 2
	})
	fmt.Println("Mapped ints:", strsInts2)
}
