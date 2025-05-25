package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator"
	"github.com/ryanbaskara/learning-go/entity"
)

func (u *Usecase) CreateUser(ctx context.Context, request *entity.CreateUserRequest) (*entity.User, error) {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) && len(validateErrs) > 0 {
			return nil, validatorError(validateErrs[0])
		}
	}

	now := time.Now()

	user := &entity.User{
		Name:        request.Name,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Status:      entity.UserStateActive,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err = u.repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
