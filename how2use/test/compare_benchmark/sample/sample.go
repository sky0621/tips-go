package sample

import "fmt"

func Before(n int) []string {
	var r []string
	for i := 0; i < n; i++ {
		r = append(r, str(i))
	}
	return r
}

func After(n int) []string {
	r := make([]string, n)
	for i := 0; i < n; i++ {
		r[i] = str(i)
	}
	return r
}

func str(n int) string {
	return fmt.Sprintf("%05d 何か", n)
}
