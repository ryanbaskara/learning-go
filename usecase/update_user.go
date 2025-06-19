package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/ryanbaskara/learning-go/entity"
)

func (u *Usecase) UpdateUser(ctx context.Context, req *entity.UpdateUserRequest) (*entity.User, error) {
	if req.ID == 0 {
		return nil, errors.New("user id should be provide")
	}

	user, err := u.GetUser(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	user.UpdatedAt = time.Now()
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.PhoneNumber != "" {
		user.PhoneNumber = req.PhoneNumber
	}

	err = u.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	u.userCacheRepository.DeleteUser(ctx, user.ID)

	err = u.userEventPublisher.PublishVerifyUser(ctx, user)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return user, nil
}
