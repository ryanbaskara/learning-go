package usecase

import (
	"context"

	"github.com/ryanbaskara/learning-go/entity"
)

type Repository interface {
	GetUser(ctx context.Context, id int64) (*entity.User, error)
	ListUsers(ctx context.Context) ([]*entity.User, error)
}

type Usecase struct {
	repository Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{
		repository: repo,
	}
}
