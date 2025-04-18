package main

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ヘルスチェック用エンドポイント
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	// パラメータ付き Hello エンドポイント例
	e.GET("/db", func(c echo.Context) error {
		message := ""
		dbConn, err := OpenDB(ctx)
		if err != nil {
			message = fmt.Sprintf("DB 接続失敗: %v", err)
		}
		defer dbConn.Close()

		message = fmt.Sprintf("DB に接続できました！")

		return c.JSON(http.StatusOK, message)
	})

	// ポート 8080 で起動
	e.Logger.Fatal(e.Start(":8080"))
}

func OpenDB(ctx context.Context) (*sql.DB, error) {
	// 必須環境変数
	user := "yuckyjuice"                          //os.Getenv("DB_USER")                  // ex: my-db-user
	pass := "yuckyjuice"                          //os.Getenv("DB_PASS")                  // ex: my-db-password
	name := "cloudsqlconnmysql01"                 //os.Getenv("DB_NAME")                  // ex: my-database
	port := "19201"                               //os.Getenv("DB_PORT")                  // ex: 3306
	inst := os.Getenv("INSTANCE_CONNECTION_NAME") // ex: project:region:instance

	// Cloud SQL Go Connector の Dialer を作成（Cloud Run／ADC 前提）
	dialer, err := cloudsqlconn.NewDialer(ctx, cloudsqlconn.WithLazyRefresh())
	if err != nil {
		return nil, fmt.Errorf("failed to create CloudSQL dialer: %w", err)
	}

	// tcp ネットワーク名を丸ごとオーバーライド
	// DSN の "@tcp(localhost:3306)" 部分がこの DialContext を経由します
	mysql.RegisterDialContext("tcp", func(ctx context.Context, addr string) (net.Conn, error) {
		if inst != "" {
			// Cloud Run 上 → addr を無視して Cloud SQL インスタンス接続
			return dialer.Dial(ctx, inst)
		}
		// ローカル → 通常の TCP ダイヤラーにフォールバック
		var d net.Dialer
		return d.DialContext(ctx, "tcp", addr)
	}) //  [oai_citation_attribution:0‡Go Packages](https://pkg.go.dev/github.com/go-sql-driver/mysql?utm_source=chatgpt.com) [oai_citation_attribution:1‡Google Cloud](https://cloud.google.com/sql/docs/mysql/samples/cloud-sql-mysql-databasesql-connect-connector?utm_source=chatgpt.com)

	// DSN は常に同じ
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?parseTime=true", user, pass, port, name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(5 * time.Minute)
	return db, nil
}
