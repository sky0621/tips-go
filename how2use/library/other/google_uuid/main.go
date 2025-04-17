package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func main() {
	v7, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	fmt.Println(v7.String())

	fmt.Println(v7.ClockSequence())

	t := v7.Time()
	fmt.Println(t)
	sec, nsec := t.UnixTime()
	t2 := time.Unix(sec, nsec)
	fmt.Println(t2)
}
