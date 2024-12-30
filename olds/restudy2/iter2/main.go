package main

import (
	"bufio"
	"io"
	"iter"
)

func lines(r io.Reader) iter.Seq[string] {
	scanner := bufio.NewScanner(r)
	return func(yield func(string) bool) {
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				break
			}
		}
	}
}

func main() {

}
