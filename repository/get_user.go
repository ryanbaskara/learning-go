package repository

import (
	"context"

	"github.com/ryanbaskara/learning-go/entity"
)

func (r *Repository) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	user := &entity.User{
		ID:   id,
		Name: "Abdul",
	}

	return user, nil
}
