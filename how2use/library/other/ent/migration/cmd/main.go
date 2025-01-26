package main

import (
	"context"
	"log"

	"github.com/sky0621/tips-go/how2use/library/other/ent/ent"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:yuckyjuice@tcp(localhost:18101)/entdb?parseTime=true")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	if err = client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
