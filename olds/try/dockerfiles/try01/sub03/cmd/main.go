package main

import (
	"fmt"

	"github.com/sky0621/tips-go/try/dockerfiles/try01/sub03"
)

func main() {
	fmt.Println("Hello, Dockerfile")
	a := &sub03.App{S: "Sample"}
	a.Print()
}
