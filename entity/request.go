package entity

type CreateUserRequest struct {
	Name        string `json:"name"         validate:"required"`
	Email       string `json:"email"        validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type UpdateUserRequest struct {
	ID          int64
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type CreateEventRequest struct {
	UserID   string                 `json:"user_id" validate:"required"`
	Type     string                 `json:"type" validate:"required"`
	Metadata map[string]interface{} `json:"metadata" validate:"required"`
}

type ListEventRequest struct {
	UserID string
}
