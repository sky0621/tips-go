package main

import (
	"database/sql"
)

func Migrate(db *sql.DB) error {
	_, err := db.Exec(createUsers)
	if err != nil {
		return err
	}
	return nil
}
