package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()

	org := os.Getenv("GRAPHQL_ORG")
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GRAPHQL_TOKEN")},
	)
	httpClient := oauth2.NewClient(ctx, src)

	client := graphql.NewClient("https://api.github.com/graphql", httpClient)

	var query struct {
		Organization struct {
			Repositories struct {
				Nodes []struct {
					Name             graphql.String
					Description      graphql.String
					DefaultBranchRef struct {
						Target struct {
							History struct {
								TotalCount graphql.Int
							} `graphql:"history(since: \"2019-01-01T00:00:00+09:00\")"`
						} `graphql:"... on Commit"`
					}
				}
			} `graphql:"repositories(first:100)"`
		} `graphql:"organization(login:$org)"`
	}
	variables := map[string]interface{}{
		"org": graphql.String(org),
	}

	err := client.Query(ctx, &query, variables)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%v", query)
}
