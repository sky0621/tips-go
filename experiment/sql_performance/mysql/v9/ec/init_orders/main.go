package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sky0621/tips-go/experiment/sql_performance/mysql/v9/ec/dsn"
)

const (
	batchSize = 100
)

func main() {
	// MySQLに接続
	db, err := sql.Open("mysql", dsn.Dsn)
	if err != nil {
		log.Fatal("DB接続エラー:", err)
	}
	defer db.Close()

	// 一括でランダムに注文と注文詳細を挿入
	insertRandomOrdersAndItems(db, 5000000)
	log.Println("データの挿入が完了しました。")
}

// ランダムに注文（orders）と注文詳細（order_items）を一括挿入
func insertRandomOrdersAndItems(db *sql.DB, orderCount int) {
	log.Println("注文と注文詳細を挿入中...")

	// 既存のユーザーIDと商品IDを取得
	userIDs := fetchExistingIDs(db, "Users")
	productIDs := fetchExistingIDs(db, "Products")
	if len(userIDs) == 0 || len(productIDs) == 0 {
		log.Fatal("データが足りません。UsersまたはProductsにデータを入れてください。")
	}

	// トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("トランザクション開始エラー:", err)
	}

	// `orders` のバッチ処理
	orderValues := []string{}
	orderArgs := []interface{}{}
	orderStatuses := []string{"pending", "completed", "canceled"}

	// 挿入する注文のデータ作成
	for i := 0; i < orderCount; i++ {
		userID := userIDs[rand.Intn(len(userIDs))]
		totalPrice := rand.Float64() * 50000
		status := orderStatuses[rand.Intn(len(orderStatuses))]
		orderDate := time.Now().AddDate(0, 0, -rand.Intn(365)).Format("2006-01-02 15:04:05")

		orderValues = append(orderValues, "(?, ?, ?, ?)")
		orderArgs = append(orderArgs, userID, totalPrice, status, orderDate)

		// バッチサイズごとに `INSERT`
		if len(orderValues) >= batchSize {
			err = executeBatchInsert(tx, "Orders", "(user_id, total_price, status, order_date)", orderValues, orderArgs)
			if err != nil {
				tx.Rollback()
				log.Fatal("Ordersバッチ挿入エラー:", err)
			}
			orderValues = []string{}
			orderArgs = []interface{}{}
		}
	}

	// 残りの `orders` を挿入
	if len(orderValues) > 0 {
		err = executeBatchInsert(tx, "Orders", "(user_id, total_price, status, order_date)", orderValues, orderArgs)
		if err != nil {
			tx.Rollback()
			log.Fatal("Ordersバッチ挿入エラー:", err)
		}
	}

	// **修正: 挿入した `order_id` を取得**
	orderIDs := fetchExistingIDs(db, "Orders")
	if len(orderIDs) == 0 {
		tx.Rollback()
		log.Fatal("Ordersの挿入後にIDが取得できませんでした。")
	}

	// `order_items` のバッチ処理
	orderItemValues := []string{}
	orderItemArgs := []interface{}{}

	for _, orderID := range orderIDs {
		itemCount := rand.Intn(5) + 1 // 各注文に1〜5個の商品を含める

		for j := 0; j < itemCount; j++ {
			productID := productIDs[rand.Intn(len(productIDs))]
			quantity := rand.Intn(5) + 1
			price := rand.Float64() * 10000

			orderItemValues = append(orderItemValues, "(?, ?, ?, ?)")
			orderItemArgs = append(orderItemArgs, orderID, productID, quantity, price)

			// バッチサイズごとに `INSERT`
			if len(orderItemValues) >= batchSize {
				err = executeBatchInsert(tx, "Order_Items", "(order_id, product_id, quantity, price)", orderItemValues, orderItemArgs)
				if err != nil {
					tx.Rollback()
					log.Fatal("Order_Itemsバッチ挿入エラー:", err)
				}
				orderItemValues = []string{}
				orderItemArgs = []interface{}{}
			}
		}
	}

	// 残りの `order_items` を挿入
	if len(orderItemValues) > 0 {
		err = executeBatchInsert(tx, "Order_Items", "(order_id, product_id, quantity, price)", orderItemValues, orderItemArgs)
		if err != nil {
			tx.Rollback()
			log.Fatal("Order_Itemsバッチ挿入エラー:", err)
		}
	}

	// コミット
	err = tx.Commit()
	if err != nil {
		log.Fatal("トランザクションコミットエラー:", err)
	}

	log.Println("注文と注文詳細の挿入完了。")
}

// 指定したテーブルから `id` を取得する関数
func fetchExistingIDs(db *sql.DB, table string) []int {
	query := fmt.Sprintf("SELECT id FROM %s", table)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("%sのID取得エラー: %v", table, err)
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Fatalf("%sのIDスキャンエラー: %v", table, err)
		}
		ids = append(ids, id)
	}
	return ids
}

// バッチ挿入を実行する関数
func executeBatchInsert(tx *sql.Tx, table string, columns string, values []string, args []interface{}) error {
	query := fmt.Sprintf("INSERT INTO %s %s VALUES %s", table, columns, strings.Join(values, ","))
	_, err := tx.Exec(query, args...)
	return err
}
