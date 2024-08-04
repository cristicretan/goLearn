package main

import (
	"log"

	"github.com/IBM/sarama"
)

type Consumer struct {
	consumer sarama.Consumer
}

// NewConsumer initializes a new consumer with Kafka broker
func NewConsumer(brokers []string) (*Consumer, error) {
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		return nil, err
	}
	return &Consumer{consumer: consumer}, nil
}

// Receive consumes messages from a Kafka topic
func (c *Consumer) Receive(topic string) error {
	partitionConsumer, err := c.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start consumer for partition: %v", err)
		return err
	}
	defer func(partitionConsumer sarama.PartitionConsumer) {
		err := partitionConsumer.Close()
		if err != nil {
			log.Fatalf("Failed to close partition consumer: %v", err)
			return
		}
	}(partitionConsumer)

	for msg := range partitionConsumer.Messages() {
		log.Printf("Received message: %s", string(msg.Value))
	}
	return nil
}
