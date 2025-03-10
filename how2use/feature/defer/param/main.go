package main

import "fmt"

func main() {
	fn()
}

func fn() {
	var str string
	defer sfn(fmt.Sprintf("defer sfn: %s", str)) // defer sfn:
	defer func() {
		sfn2(fmt.Sprintf("defer func sfn2: %s", str)) // defer func sfn2: hello world
	}()
	str = "hello world"
	sfn(fmt.Sprintf("end: %s", str)) // end: hello world
}

func sfn(s string) {
	fmt.Println(s)
}

func sfn2(s string) {
	fmt.Println(s)
}
