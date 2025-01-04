package migration

import (
	"github.com/uptrace/bun"
)

func Migrate(db *bun.DB) error {
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

	return nil
}
