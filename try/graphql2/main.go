package main

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
)

const target = "https://api.github.com/graphql"

func main() {
	client := graphql.NewClient(target)

	// make a request
	req := graphql.NewRequest(`
query { 
  viewer { 
    login
  }
}
`)

	// set any variables
	//req.Var("key", "value")

	// set header fields
	req.Header.Set("Cache-Control", "no-cache")

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData interface{}
	if err := client.Run(ctx, req, respData); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", respData)
}

type response struct {
	Name  string
	Items struct {
		Records []struct {
			Title string
		}
	}
}
