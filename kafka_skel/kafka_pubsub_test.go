/*
* @Author: Cristi Cretan
* @Date:   2024-08-03 19:27:18
* @Last Modified by:   Cristi Cretan
* @Last Modified time: 2024-08-03 19:27:32
 */
package main

import (
	"testing"
)

func TestOrderPlacedNotification(t *testing.T) {
	brokers := []string{"localhost:9092"}
	producer, err := NewProducer(brokers)
	if err != nil {
		t.Fatalf("Failed to create producer: %v", err)
	}

	consumer, err := NewConsumer(brokers)
	if err != nil {
		t.Fatalf("Failed to create consumer: %v", err)
	}

	topic := "order-topic"
	order := Order{
		orderID:    1001,
		customerID: 501,
		items:      make([]Item, 2),
	}
	order.items[0] = Item{
		1,
		2,
	}
	order.items[1] = Item{
		2,
		1,
	}
	// Produce an "Order Placed" message
	err = producer.Send(topic, "Order Placed: "+order.String())
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Implement verification logic for expected consumer behavior
	// This typically involves setting up consumers for each subsystem and validating received messages.
	// Note: Due to the asynchronous nature of Kafka, synchronization mechanisms or mocking may be required.
	go func() {
		err = consumer.Receive(topic)
		if err != nil {
			t.Errorf("Failed to receive message: %v", err)
			return
		}
	}()

}
