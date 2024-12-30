package main

import "fmt"

func main() {
	r1 := addAny(3, 2.5)
	fmt.Println(r1)

	// Vを推論できません
	//r2 := addAny(7, "a")

	var u uint64 = 23
	r3 := addInteger(u, u)
	fmt.Println(r3)

	//var i8 int8 = int8(9)
	// i8を型intとして使用できません
	//r4 := addInteger(i, i8)
	//fmt.Println(r4)

	r5 := IntsOnly[MyInt]{33, 66, 99}
	fmt.Println(r5)
}

func addAny[V int | float64 | string](a, b V) V {
	return a + b
}

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func addInteger[T Integer](a, b T) T {
	return a + b
}

type IntOnly interface{ ~int }
type IntsOnly[T IntOnly] []T
type MyInt int
