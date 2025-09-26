package config

import (
	"github.com/ryanbaskara/learning-go/eventpublisher"
	"github.com/ryanbaskara/learning-go/repository/cache"
	userrepo "github.com/ryanbaskara/learning-go/repository/mysql/user"
	"github.com/ryanbaskara/learning-go/rpc/server"
	"github.com/ryanbaskara/learning-go/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	Server  *grpc.Server
	RPCHost string
}

func (g *GRPCServer) Stop() error {
	g.Server.GracefulStop()

	return nil
}

func NewGRPCServer() (*GRPCServer, error) {
	config, err := loadRpcConfig()
	if err != nil {
		return nil, err
	}

	resource, err := initCommonResourceRPC(&config)
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer()

	err = initializeModuleRPC(resource, &config, server)
	if err != nil {
		return nil, err
	}

	// enable reflection
	reflection.Register(server)

	return &GRPCServer{
		Server:  server,
		RPCHost: config.RPCHost,
	}, nil
}

func initCommonResourceRPC(cfg *RPCConfig) (*Resource, error) {
	mysqlDB := newMysqlDatabase(&cfg.DatabaseConfig)
	redis := newRedis(&cfg.RedisConfig)
	kafkaProducer := newKafkaProducer(&cfg.KafkaConfig)

	return &Resource{
		Database:      mysqlDB,
		Redis:         redis,
		KafkaProducer: kafkaProducer,
	}, nil
}

func initializeModuleRPC(resource *Resource, cfg *RPCConfig, s *grpc.Server) error {
	userRepo := userrepo.NewUserRepository(resource.Database)
	userEventPublisher := eventpublisher.NewUserEventPublisher(resource.KafkaProducer, cfg.KafkaConfig.EventVerifyUserJobTopic)
	userCacheRepo := cache.NewUserCacheRepo(resource.Redis)
	usecase := usecase.NewUsecase(userRepo, userCacheRepo, userEventPublisher)

	serverCfg := &server.ServerConfig{
		Usecase: usecase,
	}
	server.RegisterRPCServer(s, serverCfg)

	return nil
}
