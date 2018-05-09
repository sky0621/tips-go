package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/sakila?charset=utf8&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS")))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var res []InformationSchema
	db.Raw("SELECT table_name FROM `information_schema`.`TABLES` WHERE `TABLE_SCHEMA` = 'sakila' ORDER BY table_name").Scan(&res)
	fmt.Println(res)

	for _, t := range res {
		var cols []Column
		db.Raw("SELECT column_name, column_default, is_nullable, data_type, column_type, column_key, extra, column_comment FROM `information_schema`.`COLUMNS` WHERE `TABLE_SCHEMA` = 'sakila' AND `TABLE_NAME` = ?", t.TableName).Scan(&cols)
		fmt.Println(cols)
	}
}

type InformationSchema struct {
	TableName string `gorm:"column:table_name"`
}

type Column struct {
	ColumnName    string `gorm:"column:column_name"`
	ColumnDefault string `gorm:"column:column_default"`
	IsNullable    string `gorm:"column:is_nullable"`
	DataType      string `gorm:"column:data_type"`
	ColumnType    string `gorm:"column:column_type"`
	ColumnKey     string `gorm:"column:column_key"`
	Extra         string `gorm:"column:extra"`
	ColumnComment string `gorm:"column:column_comment"`
}
