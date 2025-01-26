package main

import (
	"context"
	"fmt"

	"github.com/sky0621/tips-go/how2use/library/other/ent/ent/todo"

	"github.com/sky0621/tips-go/how2use/library/other/ent/ent"
	"github.com/sky0621/tips-go/how2use/library/other/ent/example/util"
)

func main() {
	util.WithDB(func(ctx context.Context, client *ent.Client) error {
		all, err := client.Todo.Query().Where(todo.HasParent()).All(ctx)
		if err != nil {
			return err
		}
		fmt.Println(all)

		todos, err := client.Todo.Query().Where(todo.HasChildren()).All(ctx)
		if err != nil {
			return err
		}
		fmt.Println(todos)
		return nil
	})
}
