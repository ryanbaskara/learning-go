package config

import (
	"fmt"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
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

func newMysqlDatabase(cfg *DatabaseConfig) *sqlx.DB {
	var sqlxDB *sqlx.DB
	var err error

	sqlxDB, err = sqlx.Connect("mysql", cfg.databaseSourceName())

	if err != nil {
		panic(err)
	}

	sqlxDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlxDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlxDB.SetConnMaxIdleTime(cfg.MaxIdleTime)
	sqlxDB.SetConnMaxLifetime(cfg.MaxLifetime)

	return sqlxDB
}

func newRedis(cfg *RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return rdb
}

func newKafkaProducer(cfg *KafkaConfig) *kafka.Producer {
	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": cfg.ProducerBrokers,
		},
	)
	if err != nil {
		panic(err)
	}

	// Channel untuk delivery report async
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return p
}
