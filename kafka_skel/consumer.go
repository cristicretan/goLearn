package main

import (
    "github.com/Shopify/sarama"
    "log"
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
func (c *Consumer) Receive(topic string) {
    partitionConsumer, err := c.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
    if err != nil {
        log.Fatalf("Failed to start consumer for partition: %v", err)
    }
    defer partitionConsumer.Close()

    for msg := range partitionConsumer.Messages() {
        log.Printf("Received message: %s", string(msg.Value))
    }
}
