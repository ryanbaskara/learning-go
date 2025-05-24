package entity

type CreateUserRequest struct {
	Name        string `json:"name"         validate:"required"`
	Email       string `json:"email"        validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}
