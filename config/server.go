package config

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ryanbaskara/learning-go/eventpublisher"
	"github.com/ryanbaskara/learning-go/handler"
	"github.com/ryanbaskara/learning-go/repository/cache"
	userrepo "github.com/ryanbaskara/learning-go/repository/mysql/user"
	"github.com/ryanbaskara/learning-go/usecase"
)

type Server struct {
	HttpServer *http.Server
}

func NewServer() (*Server, error) {
	cfg, err := loadServerConfig()
	if err != nil {
		return nil, err
	}

	mysqlDB := newMysqlDatabase(&cfg.DatabaseConfig)
	redis := newRedis(&cfg.RedisConfig)
	kafkaProducer := newKafkaProducer(&cfg.KafkaConfig)

	userRepo := userrepo.NewUserRepository(mysqlDB)
	userEventPublisher := eventpublisher.NewUserEventPublisher(kafkaProducer, cfg.KafkaConfig.EventVerifyUserJobTopic)
	userCacheRepo := cache.NewUserCacheRepo(redis)
	usecase := usecase.NewUsecase(userRepo, userCacheRepo, userEventPublisher)
	handler := handler.NewHandler(usecase)

	httpServer := &http.Server{
		Addr:         cfg.ServerHost,
		Handler:      handler.RegisterHandler(),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	return &Server{
		HttpServer: httpServer,
	}, nil
}
