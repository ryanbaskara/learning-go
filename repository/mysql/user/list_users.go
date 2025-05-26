package user

import (
	"context"

	"github.com/ryanbaskara/learning-go/entity"
)

func (r *UserRepository) ListUsers(ctx context.Context) ([]*entity.User, error) {
	users := []*entity.User{
		{
			ID:   1,
			Name: "Abdul",
		},
		{
			ID:   2,
			Name: "Rozak",
		},
	}

	return users, nil
}
