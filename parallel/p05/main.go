package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	mc := &MyCtx{Name: "Taro"}

	wvCtx := context.WithValue(context.Background(), "ckey", mc)

	no := 0
	for no < 10 {
		no = no + 1
		go goro(wvCtx)
	}

	time.Sleep(1 * time.Second)

	mc.Name = "Jiro"

	time.Sleep(5 * time.Second)

	mc.Name = "Saburo"

	time.Sleep(5 * time.Second)
}

type MyCtx struct {
	Name string
}

func goro(c context.Context) {
	for {
		fmt.Println(c.Value("ckey"))

		time.Sleep(3 * time.Second)
	}
}
