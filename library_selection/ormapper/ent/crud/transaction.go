package crud

import (
	"context"
	"fmt"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/ent/ent"
)

func Transaction(ctx context.Context, client *ent.Client) {
	if err := WithTx(ctx, client, func(tx *ent.Tx) error {
		post, err := tx.Post.UpdateOneID(1).SetTitle("Second Post").Save(ctx)
		if err != nil {
			return err
		}
		fmt.Println(post)
		//if post.ID == 1 {
		//	return errors.New("ERR")
		//}

		_, err = tx.Post.UpdateOneID(1).SetTitle("Third Post").Save(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			if err := tx.Rollback(); err != nil {
				fmt.Println("rollback err:", err)
				return
			}
			panic(v)
		}
	}()

	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rollback err: %v, rollback transaction: %v", err, rerr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit err: %v", err)
	}

	return nil
}
