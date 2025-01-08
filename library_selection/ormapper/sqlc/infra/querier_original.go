package infra

import "context"

type QuerierOriginal interface {
	ListUsersByIDs2(ctx context.Context, ids []int64) ([]User, error)
}
