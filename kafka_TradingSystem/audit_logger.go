package main

import (
    "github.com/Shopify/sarama"
    "log"
)

type AuditLogger struct {
    consumer sarama.Consumer
}

func NewAuditLogger(brokers []string) (*AuditLogger, error) {
    consumer, err := sarama.NewConsumer(brokers, nil)
    if err != nil {
        return nil, err
    }
    return &AuditLogger{consumer: consumer}, nil
}

func (c *AuditLogger) LogActivities() {
    // TODO: Implement logging of all messages
}

func main() {
    brokers := []string{"localhost:9092"}
    logger, err := NewAuditLogger(brokers)
    if err != nil {
        log.Fatalf("Failed to start logger: %v", err)
    }
    logger.LogActivities()
}
