package eventpublisher

import (
	"context"
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ryanbaskara/learning-go/entity"
)

type UserEventPublisher struct {
	Producer *kafka.Producer
	Topic    string
}

func NewUserEventPublisher(producer *kafka.Producer, topic string) *UserEventPublisher {
	return &UserEventPublisher{
		Producer: producer,
		Topic:    topic,
	}
}

func (p *UserEventPublisher) PublishVerifyUser(ctx context.Context, user *entity.User) error {
	payload := &entity.EventUser{
		ID:        user.ID,
		Status:    user.Status,
		EventTime: time.Now(),
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.Topic, Partition: kafka.PartitionAny},
		Value:          b,
	}, nil)

	return nil
}
