package server

import (
	pb "github.com/ryanbaskara/learning-go/protobuf/go/products/v1/public"
	"github.com/ryanbaskara/learning-go/rpc/service"
	"google.golang.org/grpc"
)

type ServerConfig struct {
	Usecase service.Usecase
}

func RegisterRPCServer(s *grpc.Server, cfg *ServerConfig) {
	publicService := service.NewPublicService(cfg.Usecase)
	pb.RegisterGreeterServer(s, publicService)
}
