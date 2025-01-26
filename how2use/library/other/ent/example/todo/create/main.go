package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sky0621/tips-go/how2use/library/other/ent/example/util"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sky0621/tips-go/how2use/library/other/ent/ent"
)

func main() {
	util.WithDB(func(ctx context.Context, client *ent.Client) error {
		save, err := client.Todo.Create().SetText("Add Example3").Save(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(save)
		return nil
	})
}
