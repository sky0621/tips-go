package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main() {
	var group errgroup.Group

	urls := []string{
		"https://www.google.com",
		"https://nonexistent.example", // 存在しないURL（エラー発生の例）
		"https://www.example.com",
		"https://nonexistent.example2", // 存在しないURL（エラー発生の例）
	}

	for _, url := range urls {
		group.Go(func() error {
			response, err := http.Get(url)
			if err != nil {
				return err
			}
			defer response.Body.Close()
			fmt.Printf("URL: %s, Status: %s\n", url, response.Status)
			return nil
		})
	}

	if err := group.Wait(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All requests succeeded")
	}
}
