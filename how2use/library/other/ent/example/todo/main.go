package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sky0621/tips-go/how2use/library/other/ent/ent"
)

func main() {
	client, err := ent.Open("mysql", "root:yuckyjuice@tcp(localhost:18101)/entdb")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	save, err := client.Todo.Create().SetText("Add Example").Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(save)
}
