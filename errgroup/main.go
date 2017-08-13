package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://dummy/",
		"http://www.google.com/",
		"http://localhost.dummy/",
		"http://example.com.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			fmt.Printf("start:[%s]\n", url)
			resp, err := http.Get(url)
			fmt.Printf("end:  [%s]\n", url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println(err)
	}
}
