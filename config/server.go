package config

import (
	"github.com/julienschmidt/httprouter"
	"github.com/ryanbaskara/learning-go/handler"
	"github.com/ryanbaskara/learning-go/repository"
	"github.com/ryanbaskara/learning-go/usecase"
)

type ServerConfig struct {
	Router *httprouter.Router
	Host   string
}

func NewServerConfig() *ServerConfig {
	repo := repository.NewRepository(nil)
	usecase := usecase.NewUsecase(repo)
	handler := handler.NewHandler(usecase)

	return &ServerConfig{
		Router: handler.RegisterHandler(),
		Host:   "localhost:7172",
	}
}
