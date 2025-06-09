package user

import (
	"context"
	"errors"

	"github.com/ryanbaskara/learning-go/entity"
)

func (r *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	query := "UPDATE users SET name = ?, email = ?, phone_number = ?, status = ?, updated_at = ? WHERE id = ?"
	query = r.db.Rebind(query)
	res, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.PhoneNumber, user.Status, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}
	rowAffected, err := res.RowsAffected()
	if err != nil {
		return nil
	}
	if rowAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}
