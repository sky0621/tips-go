package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getStdin()
	fmt.Println(lines)
	fmt.Println(len(lines))

	if len(lines) != 3 {
		return
	}

	h, _ := strconv.Atoi(lines[0])
	a, _ := strconv.Atoi(lines[1])
	b, _ := strconv.Atoi(lines[2])

	if a <= b {
		fmt.Println("NO")
		return
	}

	n := 1
	for {
		h = h - a + b
		if h <= 0 {
			fmt.Println("YES")
			fmt.Println(n)
			return
		}
		n++
	}

}

func getStdin() []string {
	stdin, _ := ioutil.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
