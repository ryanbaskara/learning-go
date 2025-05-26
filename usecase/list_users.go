package usecase

import (
	"context"

	"github.com/ryanbaskara/learning-go/entity"
)

func (u *Usecase) ListUsers(ctx context.Context) ([]*entity.User, error) {
	user, err := u.userRepository.ListUsers(ctx)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return user, nil
}
