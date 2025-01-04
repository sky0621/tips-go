package crud

import (
	"context"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/ent/ent"
)

func Update(ctx context.Context, client *ent.Client) {
	_, err := client.User.UpdateOneID(1).SetName("Alice Updated").Save(ctx)
	if err != nil {
		log.Fatalf("failed updating record: %v", err)
	}
}
