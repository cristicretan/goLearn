package main

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestProduceConsume(t *testing.T) {
	broker := GetBrokerInstance()
	consumerID := "consumer 1"

	broker.RegisterConsumer(consumerID)
	// Test producing a message
	err := broker.Produce("Order Placed: {order_id: 123, customer_id: 456, items: [{id: 1, qty: 2}, {id: 2, qty: 1}]}")
	if err != nil {
		t.Errorf("Error produce: %v", err)
	}
	if len(broker.messages) != 1 {
		t.Errorf("Expected 1 message, got %d", len(broker.messages))
	}

	// Test consuming a message
	msg, err := broker.Consume()
	if err != nil {
		t.Error(err)
	}
	if msg.Content != "Order Placed: {order_id: 123, customer_id: 456, items: [{id: 1, qty: 2}, {id: 2, qty: 1}]}" {
		t.Errorf("Expected message 'Order Placed', got %s", msg.Content)
	}

	// Test empty queue handling
	_, err = broker.Consume()

	if err.Error() != "no messages to consume" {
		t.Errorf("Expected no messages to consume, got %s", err.Error())
	}
}

func TestConcurrentProduce(t *testing.T) {
	broker := GetBrokerInstance()
	var wg sync.WaitGroup

	numMessages := 10
	messages := make([]string, numMessages)

	// Prepare messages
	for i := 0; i < numMessages; i++ {
		messages[i] = "Order Placed: {order_id: " + strconv.Itoa(i) + ", customer_id: " +
			strconv.Itoa(i+1000) + ", items: [{id: " + strconv.Itoa(i%10) + ", qty: " + strconv.Itoa((i%5)+1) + "}]}"
	}

	// Start multiple goroutines to produce messages concurrently
	for i := 0; i < numMessages; i++ {
		wg.Add(1)
		go func(msg string) {
			defer wg.Done()
			if err := broker.Produce(msg); err != nil {
				t.Errorf("Failed to produce message: %s", err)
			}
		}(messages[i])
	}

	wg.Wait()

	if len(broker.messages) != numMessages {
		t.Errorf("Expected %d, got %d", numMessages, len(broker.messages))
	}

	for i := 0; i < numMessages; i++ {
		fmt.Printf("message %s\n", broker.messages[i].Content)
	}
}

func TestConcurrentConsume(t *testing.T) {
	broker := NewMyKafka()
	var wg sync.WaitGroup

	// Number of messages and consumers
	numMessages := 10
	numConsumers := 3

	for i := 0; i < numConsumers; i++ {
		broker.RegisterConsumer("consumer " + strconv.Itoa(i))
	}

	// Produce messages
	for i := 0; i < numMessages; i++ {
		msg := fmt.Sprintf("Order Placed: {order_id: %d\n}", i)
		if err := broker.Produce(msg); err != nil {
			t.Errorf("Failed to produce message: %s", err)
		}
	}

	// Map to track which messages have been consumed
	consumedMessages := make(map[int]bool)
	var mu sync.Mutex

	// Function for each consumer to consume messages
	consumeMessages := func(consumerID string) {
		defer wg.Done()
		for {
			// Attempt to consume a message
			msg, err := broker.Consume()
			if err != nil {
				if err.Error() == "no messages to consume" {
					return // No more messages
				}
				t.Errorf("Consumer %s failed to consume message: %s", consumerID, msg.Content)
			}

			// Register the consumed message
			mu.Lock()
			consumedMessages[msg.ID] = true
			mu.Unlock()
		}
	}

	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consumeMessages(fmt.Sprintf("%d", i))
	}
	wg.Wait()

	if len(consumedMessages) != numMessages {
		t.Errorf("Expected %d, got %d", numMessages, len(consumedMessages))
	}
}

func TestProducerSend(t *testing.T) {
	broker := GetBrokerInstance()
	producer := NewProducer()

	producer.Send("Order Placed: {order_id: 1001, customer_id: 501, items: [{id: 1, qty: 2}]}")

	if len(broker.messages) != 1 {
		t.Errorf("Expected 1 message, got %d", len(broker.messages))
	}
}

func TestConsumerReceive(t *testing.T) {
	broker := GetBrokerInstance()
	consumer := NewConsumer("consumer1")

	err := broker.Produce("Order Placed: {order_id: 1001, customer_id: 501, items: [{id: 2, qty: 3}]}")
	if err != nil {
		t.Fatalf("Failed to produce message: %s", err)
	}

	message, err := consumer.Receive()
	if err != nil {
		t.Fatalf("Error receiving message: %v", err)
	}

	expectedMessage := "Order Placed: {order_id: 1001, customer_id: 501, items: [{id: 2, qty: 3}]}"
	if message != expectedMessage {
		t.Errorf("Expected message '%s', got '%s'", expectedMessage, message)
	}
}

//
//func TestLoadBalancing(t *testing.T) {
//	broker := NewMyKafka()
//	broker.RegisterConsumer("Consumer1")
//	broker.RegisterConsumer("Consumer2")
//	broker.Produce("Order Placed: {order_id: 125, customer_id: 458, items: [{id: 4, qty: 5}]}")
//
//	// Implement logic to check if the message is distributed correctly
//}
