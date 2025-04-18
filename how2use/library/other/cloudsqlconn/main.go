package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	dbConn, err := OpenDB(ctx)
	if err != nil {
		log.Fatalf("DB 接続失敗: %v", err)
	}
	defer dbConn.Close()

	// 以降 dbConn.Query / dbConn.Exec で共通コード
	log.Println("DB に接続できました！")
}
