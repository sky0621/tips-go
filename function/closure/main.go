package main

import "fmt"

func main() {
	d := multi(3)
	fmt.Println(d(4))

	res := multi(5)(10)
	fmt.Println(res)
}

type double func(i int) int

func multi(times int) double {
	return func(i int) int {
		return times * i
	}
}
