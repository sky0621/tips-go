package crud

import (
	"context"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/ent/ent"
)

// Delete ... 物理削除
func Delete(ctx context.Context, client *ent.Client) {
	if err := client.User.DeleteOneID(1).Exec(ctx); err != nil {
		log.Fatalf("failed updating record: %v", err)
	}
}
