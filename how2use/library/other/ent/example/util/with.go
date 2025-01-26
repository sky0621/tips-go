package util

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sky0621/tips-go/how2use/library/other/ent/ent"
)

func WithDB(fn func(ctx context.Context, client *ent.Client) error) {
	client, err := ent.Open("mysql", "root:yuckyjuice@tcp(localhost:18101)/entdb?parseTime=true")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	if err := fn(ctx, client); err != nil {
		log.Fatalf("failed executing function: %v", err)
	}
}
