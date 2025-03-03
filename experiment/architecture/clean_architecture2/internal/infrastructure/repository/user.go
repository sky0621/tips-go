package repository

import (
	"context"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture2/internal/domain"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture2/internal/infrastructure/db"
)

type UserRepositoryImpl struct {
	queries *db.Queries
}

func (u UserRepositoryImpl) FindUserByID(ctx context.Context, id int64) (domain.User, error) {
	user, err := u.queries.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func NewUser(queries *db.Queries) domain.UserRepository {
	return &UserRepositoryImpl{queries: queries}
}
