package repository

import (
	"context"

	"github.com/ryanbaskara/learning-go/entity"
)

func (r *Repository) CreateUser(ctx context.Context, user *entity.User) error {
	q := "INSERT INTO users (name, email, phone_number, status, created_at, updated_at) "
	q += "VALUES (:name, :email, :phone_number, :status, :created_at, :updated_at)"

	res, err := r.db.NamedExecContext(ctx, q, user)
	if err != nil {
		return err
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = lastInsertId

	return nil
}
