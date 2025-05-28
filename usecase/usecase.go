package usecase

import (
	"context"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/ryanbaskara/learning-go/entity"
)

//go:generate mockgen -package=mock_usecase -source=usecase.go -destination=mocks/usecase.go

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, id int64) (*entity.User, error)
	ListUsers(ctx context.Context) ([]*entity.User, error)
}

type UserCacheRepository interface {
	GetUser(ctx context.Context, id int64) (*entity.User, error)
	SetUser(ctx context.Context, user *entity.User) error
}

type Usecase struct {
	userRepository      UserRepository
	userCacheRepository UserCacheRepository
}

func NewUsecase(
	userRepository UserRepository,
	userCacheRepository UserCacheRepository,
) *Usecase {
	return &Usecase{
		userRepository:      userRepository,
		userCacheRepository: userCacheRepository,
	}
}

func validatorError(fe validator.FieldError) error {
	switch fe.Tag() {
	case "required":
		return fmt.Errorf("%s is required", fe.Field())
	default:
		return fmt.Errorf("%s is not valid", fe.Field())
	}
}
