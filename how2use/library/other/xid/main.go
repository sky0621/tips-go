package main

import (
	"fmt"
	"github.com/rs/xid"
)

func main() {
	guid := xid.New()
	// f.e. d016ertijnjgc7bt70lg
	fmt.Println(guid.String())
}
