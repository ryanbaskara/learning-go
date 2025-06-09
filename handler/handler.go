package handler

import (
	"context"

	"github.com/julienschmidt/httprouter"
	"github.com/ryanbaskara/learning-go/entity"
)

type UseCase interface {
	CreateUser(ctx context.Context, request *entity.CreateUserRequest) (*entity.User, error)
	GetUser(ctx context.Context, id int64) (*entity.User, error)
	ListUsers(ctx context.Context) ([]*entity.User, error)
	UpdateUser(ctx context.Context, req *entity.UpdateUserRequest) (*entity.User, error)
}

type Handler struct {
	UseCase UseCase
}

func NewHandler(useCase UseCase) *Handler {
	return &Handler{
		UseCase: useCase,
	}
}

func (h *Handler) RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/health", h.Health)

	router.POST("/users", h.CreateUser)
	router.GET("/users", h.ListUsers)
	router.GET("/users/:user_id", h.GetUser)
	router.PATCH("/users/:user_id", h.UpdateUser)

	return router
}
