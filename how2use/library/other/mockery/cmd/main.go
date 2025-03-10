package main

import (
	"fmt"

	"github.com/sky0621/tips-go/how2use/library/other/mockery/mocks"
)

func main() {
	fmt.Println("mockery")
	t := &mocks.ProgrammingLanguage{}
	fmt.Println(t.GetName())
}
