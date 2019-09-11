package main

import "testing"

func TestFactorial(t *testing.T) {
	var expect uint64 = 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	actual := factorial(10, 1)
	if expect != actual {
		t.Errorf("Expect: %d, but got %d", expect, actual)
	}
}
