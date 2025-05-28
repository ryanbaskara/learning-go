package usecase

import (
	"context"

	"github.com/ryanbaskara/learning-go/entity"
)

func (u *Usecase) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	user, err := u.userCacheRepository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	if user != nil {
		user.Source = "redis"
		return user, nil
	}

	user, err = u.userRepository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	u.userCacheRepository.SetUser(ctx, user)
	user.Source = "mysql"

	return user, nil
}
