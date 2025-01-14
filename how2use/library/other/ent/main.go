package main

import (
	"context"
	"log"

	"github.com/sky0621/tips-go/how2use/library/other/ent/ent"
	"github.com/sky0621/tips-go/how2use/library/other/ent/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "test01:yuckyjuice@tcp(localhost:18001)/localmysqlent01")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// マイグレーションの実行
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
