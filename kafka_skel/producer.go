/*
* @Author: Cristi Cretan
* @Date:   2024-08-03 19:26:31
* @Last Modified by:   Cristi Cretan
* @Last Modified time: 2024-08-03 19:26:43
*/
package main

import (
    "github.com/Shopify/sarama"
    "log"
)

type Producer struct {
    producer sarama.SyncProducer
}

// NewProducer initializes a new producer with Kafka broker
func NewProducer(brokers []string) (*Producer, error) {
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true

    producer, err := sarama.NewSyncProducer(brokers, config)
    if err != nil {
        return nil, err
    }
    return &Producer{producer: producer}, nil
}

// Send sends a message to a Kafka topic
func (p *Producer) Send(topic, message string) error {
    msg := &sarama.ProducerMessage{
        Topic: topic,
        Value: sarama.StringEncoder(message),
    }
    _, _, err := p.producer.SendMessage(msg)
    if err != nil {
        return err
    }
    log.Printf("Message sent to topic %s: %s", topic, message)
    return nil
}
