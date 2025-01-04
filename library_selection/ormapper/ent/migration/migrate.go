package migration

import (
	"context"

	"github.com/sky0621/tips-go/library_selection/ormapper/ent/ent"
)

func Migrate(client *ent.Client) error {
	return client.Schema.Create(context.Background())
}
