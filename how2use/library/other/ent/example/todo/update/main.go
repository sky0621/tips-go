package main

import (
	"context"

	"github.com/sky0621/tips-go/how2use/library/other/ent/example/util"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sky0621/tips-go/how2use/library/other/ent/ent"
)

func main() {
	util.WithDB(func(ctx context.Context, client *ent.Client) error {
		todos, err := client.Todo.Query().All(ctx)
		if err != nil {
			return err
		}
		if err := todos[1].Update().SetParent(todos[0]).Exec(ctx); err != nil {
			return err
		}
		return nil
	})
}
