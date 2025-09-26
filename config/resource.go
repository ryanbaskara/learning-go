package config

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Resource struct {
	Database      *sqlx.DB
	Redis         *redis.Client
	KafkaProducer *kafka.Producer
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
