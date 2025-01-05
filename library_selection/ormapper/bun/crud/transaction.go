package crud

import (
	"context"
	"log"
	"time"

	"github.com/sky0621/tips-go/library_selection/ormapper/bun/models"

	"github.com/uptrace/bun"
)

func Transaction(ctx context.Context, db *bun.DB) {
	if err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		u := &models.User{
			Name:      "Bob",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		_, err := tx.NewInsert().Model(u).Exec(ctx)
		if err != nil {
			return err
		}
		//if u.Name == "Bob" {
		//	return errors.New("ERR")
		//}

		u2 := &models.User{
			Name:      "Ken",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		_, err = tx.NewInsert().Model(u2).Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
