package db

import (
    "fmt"
	"os"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer

func InitKafkaProducer() {
    var err error
    producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
		"sasl.username": os.Getenv("KAFKA_USERNAME"),
		"sasl.password": os.Getenv("KAFKA_PASSWORD"),
	})
    if err != nil {
        fmt.Printf("Failed to create producer: %s\n", err)
    }
}

func PublishEvent(topic string, message string) {
    err := producer.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value:          []byte(message),
    }, nil)
    if err != nil {
        fmt.Printf("Error producing message to %s: %s\n", topic, err)
    } else {
        fmt.Printf("Message successfully produced to %s!\n", topic)
    }
}

