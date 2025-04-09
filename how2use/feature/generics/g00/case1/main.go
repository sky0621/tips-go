package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	strs := []string{"a", "b", "c", "d", "e"}

	/*
	 * Before generics
	 */
	fmt.Println(containsInt(nums, 3))
	fmt.Println(containsInt(nums, 6))

	fmt.Println(containsString(strs, "c"))
	fmt.Println(containsString(strs, "f"))

	/*
	 * After generics
	 */
	fmt.Println(contains(nums, 3))
	fmt.Println(contains(nums, 6))

	fmt.Println(contains(strs, "c"))
	fmt.Println(contains(strs, "f"))
}
