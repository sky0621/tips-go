package migration

import (
	"database/sql"
)

func Migrate(db *sql.DB) error {
	ddls := []string{
		createUsers,
		createPosts,
		createComments,
		createDepartments,
		createEmployees,
		createSchool,
		createGrade,
		createClass,
		createStudent,
	}
	for _, ddl := range ddls {
		_, err := db.Exec(ddl)
		if err != nil {
			return err
		}
	}
	return nil
}
