package main

import "fmt"

func main() {
	list := []string{"go", "java", "scala", "ruby", "python", "c"}
	printEl(list)
}

func printEl(list []string) {
	if list == nil || len(list) == 0 {
		return
	}
	fmt.Println(list[0])
	if len(list) == 1 {
		return
	}
	printEl(list[1:])
}
