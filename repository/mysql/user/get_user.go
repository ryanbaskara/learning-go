package user

import (
	"context"

	"github.com/ryanbaskara/learning-go/entity"
)

func (r *UserRepository) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	q := "SELECT id, name, email, phone_number, status, created_at, updated_at FROM users "
	q += "WHERE id = ?"
	q = r.db.Rebind(q)

	var user entity.User
	if err := r.db.GetContext(ctx, &user, q, id); err != nil {
		return nil, err
	}

	return &user, nil
}
