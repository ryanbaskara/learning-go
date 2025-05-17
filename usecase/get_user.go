package usecase

import (
	"context"

	"github.com/ryanbaskara/learning-go/entity"
)

func (u *Usecase) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	user, err := u.repository.GetUser(ctx, id)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return user, nil
}
