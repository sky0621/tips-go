package main

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	"github.com/sky0621/tips-go/experiment/sql_performance/mysql/v9/ec/dsn"

	"github.com/bxcodec/faker/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// MySQLに接続
	db, err := sql.Open("mysql", dsn.Dsn)
	if err != nil {
		log.Fatal("DB接続エラー:", err)
	}
	defer db.Close()

	count := 1000000
	insertUsers(db, count)
	insertProducts(db, count)
	insertOrders(db, count)
	insertOrderItems(db, count)
	insertPayments(db, count)

	log.Println("データの挿入が完了しました。")
}

func insertUsers(db *sql.DB, count int) {
	log.Println("ユーザーを挿入中...")
	stmt, _ := db.Prepare("INSERT INTO Users (name, email, password_hash) VALUES (?, ?, ?)")
	for i := 0; i < count; i++ {
		_, err := stmt.Exec(faker.Name(), faker.Email(), faker.Password())
		if err != nil {
			log.Println("Users挿入エラー:", err)
		}
	}
	log.Println("ユーザー挿入完了。")
}

func insertProducts(db *sql.DB, count int) {
	log.Println("商品を挿入中...")
	stmt, _ := db.Prepare("INSERT INTO Products (name, description, price, stock_quantity) VALUES (?, ?, ?, ?)")
	for i := 0; i < count; i++ {
		_, err := stmt.Exec(faker.Word(), faker.Sentence(), rand.Float64()*10000, rand.Intn(100))
		if err != nil {
			log.Println("Products挿入エラー:", err)
		}
	}
	log.Println("商品挿入完了。")
}

func insertOrders(db *sql.DB, count int) {
	log.Println("注文を挿入中...")
	stmt, _ := db.Prepare("INSERT INTO Orders (user_id, total_price, status, order_date) VALUES (?, ?, ?, ?)")
	for i := 0; i < count; i++ {
		userID := rand.Intn(count) + 1
		totalPrice := rand.Float64() * 50000
		statuses := []string{"pending", "completed", "canceled"}
		status := statuses[rand.Intn(len(statuses))]
		orderDate := time.Now().AddDate(0, 0, -rand.Intn(365)).Format("2006-01-02 15:04:05")

		_, err := stmt.Exec(userID, totalPrice, status, orderDate)
		if err != nil {
			log.Println("Orders挿入エラー:", err)
		}
	}
	log.Println("注文挿入完了。")
}

func insertOrderItems(db *sql.DB, count int) {
	log.Println("注文詳細を挿入中...")
	stmt, _ := db.Prepare("INSERT INTO Order_Items (order_id, product_id, quantity, price) VALUES (?, ?, ?, ?)")
	for i := 0; i < count; i++ {
		orderID := rand.Intn(count) + 1
		productID := rand.Intn(count) + 1
		quantity := rand.Intn(5) + 1
		price := rand.Float64() * 10000

		_, err := stmt.Exec(orderID, productID, quantity, price)
		if err != nil {
			log.Println("Order_Items挿入エラー:", err)
		}
	}
	log.Println("注文詳細挿入完了。")
}

func insertPayments(db *sql.DB, count int) {
	log.Println("支払い情報を挿入中...")
	stmt, _ := db.Prepare("INSERT INTO Payments (order_id, payment_method, payment_status) VALUES (?, ?, ?)")
	for i := 0; i < count; i++ {
		orderID := rand.Intn(count) + 1
		methods := []string{"credit_card", "paypal", "bank_transfer"}
		paymentMethod := methods[rand.Intn(len(methods))]
		statuses := []string{"success", "failed", "pending"}
		paymentStatus := statuses[rand.Intn(len(statuses))]

		_, err := stmt.Exec(orderID, paymentMethod, paymentStatus)
		if err != nil {
			log.Println("Payments挿入エラー:", err)
		}
	}
	log.Println("支払い情報挿入完了。")
}
