package main

import "fmt"

func main() {
	printInt(3)

	mi := MyInt(8)
	printInt(mi)
}

func printInt[T ~int](x T) {
	fmt.Println(x)
}

type MyInt int
