package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func main() {
	// エコー対象の言語は５つ
	langs := []string{
		"Go",
		"    Java",
		"         C++",
		"              PHP",
		"                  Ruby"}

	ctx, cancelFunc := context.WithCancel(context.Background())

	for _, lang := range langs {
		go echoLangName(ctx, lang)
	}

	rec := newSignalReceiver(cancelFunc)
	rec.wait()

	time.Sleep(3 * time.Second) // 本来はWaitGroup使う
}

// １秒おきにnameを出力
func echoLangName(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("***** stop echo " + strings.TrimSpace(name) + " *****")
			return
		default:
			fmt.Println(name)
			time.Sleep(1 * time.Second)
			continue
		}
	}
}
