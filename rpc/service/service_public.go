package service

import (
	"context"
	"log"

	"github.com/ryanbaskara/learning-go/entity"
	pb "github.com/ryanbaskara/learning-go/protobuf/go/products/v1/public"
)

type Usecase interface {
	CreateUser(ctx context.Context, request *entity.CreateUserRequest) (*entity.User, error)
	GetUser(ctx context.Context, id int64) (*entity.User, error)
	ListUsers(ctx context.Context) ([]*entity.User, error)
	UpdateUser(ctx context.Context, req *entity.UpdateUserRequest) (*entity.User, error)
}

type PublicService struct {
	pb.UnimplementedGreeterServer
	Usecase Usecase
}

func NewPublicService(Usecase Usecase) *PublicService {
	return &PublicService{
		Usecase: Usecase,
	}
}

func (s *PublicService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
