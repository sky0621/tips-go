package main

import "fmt"

type MyStr string

func formatString(s string) string {
	return fmt.Sprintf("str:%s", s)
}

func formatStringImpl[V ~string](s V) V {
	return V(fmt.Sprintf("str.impl:%s", s))
}
