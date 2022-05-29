package main

import "fmt"

func main() {
	//sample1();
	//sample2()
	//sample3()
	sample4()
}

func sample4() {
	fmt.Println(formatString("one"))
	var ms MyStr
	ms = "MyStr"
	fmt.Println(ms)
	fmt.Println(formatStringImpl(ms))
}

func sample3() {
	type Favorite struct {
		Food string
	}
	var f Favorite
	if err := JsonDecode([]byte(`{"food": "中本"}`), &f); err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(f)
}

func sample2() {
	var v Vector[int]
	v.Push(3)
	fmt.Println(v)
	v.Push(25)
	fmt.Println(v)
}

func sample1() {
	fmt.Println(addInts(1, 2))
	fmt.Println(addFloats(2.2, 3.3))
	fmt.Println(addString("a", "b"))

	fmt.Println(addAny[int](5, 6))
	fmt.Println(addAny[string]("7", "8"))

	o := &One{
		ID:   122,
		Name: "OneOne",
	}

	t := &Two{ID: 5656, IsAdmin: true}

	something(*o)
	something(*t)
}
