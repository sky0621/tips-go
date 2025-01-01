package main

import util "github.com/sky0621/tips-go/knowledge/goroutine_throttling"

func main() {
	util.ExecMain(func() {
		for i := 0; i < 50; i++ {
			util.Download(i)
		}
	})
}
