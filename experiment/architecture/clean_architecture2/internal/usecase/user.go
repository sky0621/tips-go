package usecase

import (
	"context"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture2/internal/domain"
)

type GetUserInputPort interface {
	GetUser(ctx context.Context, req GetUserRequest, presenter GetUserOutputPort) error
}

type getUserInteractor struct {
	repo domain.UserRepository
}

func (g getUserInteractor) GetUser(ctx context.Context, req GetUserRequest, presenter GetUserOutputPort) error {
	user, err := g.repo.FindUserByID(ctx, req.UserID)
	if err != nil {
		return err
	}
	return presenter.Present(GetUserResponse{User: user})
}

func NewGetUserInteractor(repo domain.UserRepository) GetUserInputPort {
	return &getUserInteractor{repo: repo}
}

type GetUserRequest struct {
	UserID int64
}

type GetUserOutputPort interface {
	Present(res GetUserResponse) error
}

type GetUserResponse struct {
	User domain.User
}
