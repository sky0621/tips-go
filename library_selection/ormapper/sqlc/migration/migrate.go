package migration

import (
	"database/sql"
)

func Migrate(db *sql.DB) error {
	_, err := db.Exec(createUsers)
	if err != nil {
		return err
	}

	_, err = db.Exec(createPosts)
	if err != nil {
		return err
	}

	_, err = db.Exec(createComments)
	if err != nil {
		return err
	}

	_, err = db.Exec(createDepartments)
	if err != nil {
		return err
	}

	_, err = db.Exec(createEmployees)
	if err != nil {
		return err
	}

	return nil
}
