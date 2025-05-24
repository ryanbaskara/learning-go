package repository

import (
	"context"
	"time"

	"github.com/ryanbaskara/learning-go/entity"
)

func (r *Repository) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	user := &entity.User{
		ID:          id,
		Name:        "Abdul",
		Email:       "aaa@bcd.com",
		PhoneNumber: "0899999",
		Status:      entity.UserStateActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return user, nil
}
