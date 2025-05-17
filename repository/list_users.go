package repository

import (
	"context"

	"github.com/ryanbaskara/learning-go/entity"
)

func (r *Repository) ListUsers(ctx context.Context) ([]*entity.User, error) {
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
