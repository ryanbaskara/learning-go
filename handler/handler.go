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

type EventUseCase interface {
	CreateEvent(ctx context.Context, req *entity.CreateEventRequest) (*entity.Event, error)
	ListEventByUserID(ctx context.Context, req *entity.ListEventRequest) ([]*entity.Event, error)
}

type Handler struct {
	UseCase      UseCase
	eventUseCase EventUseCase
}

func NewHandler(
	useCase UseCase,
	eventUseCase EventUseCase,
) *Handler {
	return &Handler{
		UseCase:      useCase,
		eventUseCase: eventUseCase,
	}
}

func (h *Handler) RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/health", h.Health)

	router.POST("/events", h.CreateEvent)
	router.GET("/events", h.ListEvents)

	router.POST("/users", h.CreateUser)
	router.GET("/users", h.ListUsers)
	router.GET("/users/:user_id", h.GetUser)
	router.PATCH("/users/:user_id", h.UpdateUser)

	return router
}
