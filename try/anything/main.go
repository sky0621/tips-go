package main

import "fmt"

func main() {
	s := &MyStruct{}
	fn(s)

	var i64 int64 = 3
	fn2(i64)

	var f64 = 7.2
	fn2(f64)

	var a A = 35
	var b B = 45
	var o O = 55
	var ab AB = 65
	//var z int = 75
	fn3(a)
	fn3(b)
	fn3(o)
	fn3(ab)
	//fn3(z)

	s2 := S[int64]{e: 333}
	var e2 int64 = 444
	fmt.Println(s2.Add(e2))
}

type S[N Number] struct {
	e N
}

func (s S[N]) Add(e N) N {
	return s.e + e
}

type MyInterface interface {
	Show()
}

type MyStruct struct{}

func (*MyStruct) Show() {
	fmt.Println("MyStruct show")
}

func fn[T MyInterface](p T) {
	p.Show()
}

type Number interface {
	int64 | float64
}

func fn2[T Number](p T) {
	fmt.Println(p)
}

type A int
type B int
type O int
type AB int

type BloodType interface {
	A | B | O | AB
}

func fn3[T BloodType](p T) {
	fmt.Println(p)
}
