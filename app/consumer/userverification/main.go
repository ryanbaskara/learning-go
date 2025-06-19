package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "users-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.SubscribeTopics([]string{"learning.verify-user-job"}, nil)

	fmt.Println("Consumer started...")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
