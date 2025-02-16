package main

import (
	"database/sql"
)

func Migrate(db *sql.DB) error {
	_, err := db.Exec(createTodos)
	if err != nil {
		return err
	}
	return nil
}
