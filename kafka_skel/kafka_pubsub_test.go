/*
* @Author: Cristi Cretan
* @Date:   2024-08-03 19:27:18
* @Last Modified by:   Cristi Cretan
* @Last Modified time: 2024-08-03 19:27:32
 */
package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type MockConsumer struct {
	messages []string
	mu       sync.Mutex
}

func newMockConsumer() *MockConsumer {
	return &MockConsumer{
		messages: []string{},
	}
}

func (m *MockConsumer) Receive(topic string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.messages = append(m.messages, fmt.Sprintf("Message for topic %s", topic))
	return nil
}

func (m *MockConsumer) GetMessages() []string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.messages
}

func TestOrderPlacedNotification(t *testing.T) {
	brokers := []string{"localhost:9092"}
	inventoryMock := newMockConsumer()
	billingMock := newMockConsumer()
	shippingMock := newMockConsumer()

	inventorySystem := newInventorySystem()
	billingSystem := newBillingSystem(1000)
	shippingSystem := newShippingSystem()

	inventorySystem.Start()
	billingSystem.Start()
	shippingSystem.Start()

	producer, err := NewProducer(brokers)
	if err != nil {
		t.Fatalf("Failed to create producer: %v", err)
	}

	order := Order{
		OrderID:    1001,
		CustomerID: 501,
		Items:      make([]Item, 2),
	}
	order.Items[0] = Item{
		1,
		-2,
	}
	order.Items[1] = Item{
		2,
		-1,
	}
	// Produce an "Order Placed" message
	err = placeOrder(producer, order)
	if err != nil {
		t.Fatalf("Failed to place order: %v", err)
		return
	}

	err = inventoryMock.Receive(inventorySystemTopic)
	if err != nil {
		t.Fatalf("Failed to receive inventory: %v", err)
		return
	}

	err = billingMock.Receive(billingSystemTopic)
	if err != nil {
		t.Fatalf("Failed to receive billing: %v", err)
		return
	}

	err = shippingMock.Receive(shippingSystemTopic)
	if err != nil {
		t.Fatalf("Failed to receive shipping: %v", err)
		return
	}

	time.Sleep(time.Second)

	if len(inventoryMock.GetMessages()) == 0 {
		t.Fatalf("Inventory system did not receive messages")
		return
	}

	if len(billingMock.GetMessages()) == 0 {
		t.Fatalf("Billing system did not receive messages")
		return
	}

	if len(shippingMock.GetMessages()) == 0 {
		t.Fatalf("Shipping system did not receive messages")
		return
	}
}
