package main

import "os"

func main() {
	fs, err := os.Stat("xxx.mp4")
	if err != nil {
		panic(err)
	}

}
