package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GRAPHQL_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := graphql.NewClient("http://localhost:5050/query", httpClient)

	var query struct {
		User struct {
			ID    graphql.String
			Name  graphql.String
			Todos []struct {
				Text graphql.String
			}
		} `graphql:"user(id:\"57e85ed87bf64eb79caebb0aee2d79dc\")"`
	}

	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%#v", query)
}
