package infra

import "context"

type QuerierOriginal interface {
	ListUsersByIDs(ctx context.Context, ids []int64) ([]User, error)
}
