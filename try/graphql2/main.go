package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

func main() {
	//org := os.Getenv("GRAPHQL_ORG")
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GRAPHQL_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := graphql.NewClient("https://api.github.com/graphql", httpClient)

	var query struct {
		Organization struct {
			Repositories struct {
				TotalCount graphql.Int
				Nodes      []struct {
					Name        graphql.String
					Description graphql.String
				}
			} `graphql:"repositories(first:100)"`
		} `graphql:"organization(login:\"\")"`
	}

	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%#v", query)
}
