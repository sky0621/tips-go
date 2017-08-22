package main

import "fmt"

func main() {
	fmt.Println(Vo(0, 0))
}

func Vo(num int, all int) int {
	if num > 10 {
		return all
	}
	return Vo(num+1, all+num)
}
