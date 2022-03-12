package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func main() {
	print("hello comsumer!")

	// Read environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		print("Error loading .env file")
	}

	var (
		bootstrapServers = os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
		kafkaTopic       = os.Getenv("KAFKA_TOPIC_MULTI_BROKER_TOPIC")
		consumerGroupID  = os.Getenv("KAFKA_CONSUMER_GROUP_ID")
	)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          consumerGroupID,
	})

	if err != nil {
		panic(err)
	}

	err = c.Subscribe(kafkaTopic, nil)
	if err != nil {
		print("Subscribe failed:", err)
	}

	defer c.Close()

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

}
