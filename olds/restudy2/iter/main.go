package main

import "time"

func iter[T any](a []T) func() (T, bool) {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, v := range a {
			time.Sleep(1 * time.Second)
			ch <- v
		}
	}()
	return func() (T, bool) {
		v, ok := <-ch
		return v, ok
	}
}

func main() {
	vv := iter([]int{1, 2, 3, 4, 5})
	for {
		v, ok := vv()
		if !ok {
			break
		}
		println(v)
	}
}
