package util

import (
	"fmt"
	"time"
)

func Download(i int) {
	url := fmt.Sprintf("https://example.com/api/users?id=%d", i)
	println(url)
	time.Sleep(2 * time.Second)
}
