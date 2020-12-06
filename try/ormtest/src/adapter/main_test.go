package adapter

import (
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	db := setupDB()
	defer teardownDB(db)

	m.Run() // go v1.15 からは os.Exit 不要らしい
}

func setupDB() *sqlx.DB {
	// MEMO: ローカルでしか使わないので、ベタ書き
	dsn := "host=localhost port=33456 dbname=tips-go-try-ormtest-testdb user=postgres password=yuckyjuice sslmode=disable"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func teardownDB(db *sqlx.DB) {
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()
	rows, err := db.Queryx(`SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = 'public' AND tablename != 'migration'`)
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Println(err)
			return
		}
		// RESTART IDENTITY ... 消去されるテーブルの列により所有されるシーケンスを自動的に再起動させます。
		// CASCADE ... 指定されたテーブル、または、CASCADEにより削除対象テーブルとされたテーブルを参照する外部キーを持つテーブルすべてを自動的に空にします。
		db.MustExec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tableName))
	}
}
