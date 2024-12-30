package main

import "fmt"

func main() {
	// 4 4 4 4 = 4
	v := 4
	n := 4
	r := 4

	c(v, n, r)
}

func c(v, n, r int) {
	for t := 0; t < n; t++ {
		fmt.Println(v)
	}
}
