/*
* @Author: Cristi Cretan
* @Date:   2024-08-03 19:27:18
* @Last Modified by:   Cristi Cretan
* @Last Modified time: 2024-08-03 19:27:32
 */
package main

import (
	"sync"
	"testing"
)

var inventorySystem *InventorySystem
var billingSystem *BillingSystem
var shippingSystem *ShippingSystem

func TestOrderPlacedNotification(t *testing.T) {
	brokers := []string{"localhost:9092"}
	inventorySystem = newInventorySystem()
	inventorySystem.Items = make([]Item, 5)
	inventorySystem.Items[0] = Item{
		ID:  1,
		Qty: 5,
	}
	inventorySystem.Items[1] = Item{
		ID:  2,
		Qty: 5,
	}
	inventorySystem.Items[2] = Item{
		ID:  3,
		Qty: 5,
	}
	inventorySystem.Items[3] = Item{
		ID:  4,
		Qty: 5,
	}
	inventorySystem.Items[4] = Item{
		ID:  5,
		Qty: 5,
	}
	billingSystem = newBillingSystem(1000)
	shippingSystem = newShippingSystem()
	shippingSystem.Customers = make([]Customer, 3)
	shippingSystem.Customers[0] = Customer{
		ID:      501,
		Address: "Nowhere is your first customer",
	}
	shippingSystem.Customers[1] = Customer{
		ID:      502,
		Address: "Somewhere is your first customer",
	}
	shippingSystem.Customers[2] = Customer{
		ID:      503,
		Address: "Some address is your first customer",
	}
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
	var wg sync.WaitGroup

	// Implement verification logic for expected consumer behavior
	// This typically involves setting up consumers for each subsystem and validating received messages.
	// Note: Due to the asynchronous nature of Kafka, synchronization mechanisms or mocking may be required.
	go func() {
		wg.Add(1)
		err := inventorySystem.Consumer.Receive(invetorySystemTopic)
		if err != nil {
			t.Errorf("Failed to consume message on invetory system: %v", err)
			return
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		err = billingSystem.Consumer.Receive(billingSystemTopic)
		if err != nil {
			t.Errorf("Failed to consume message on billing system: %v", err)
			return
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		err = shippingSystem.Consumer.Receive(shippingSystemTopic)
		if err != nil {
			t.Errorf("Failed to consume message on shipping system: %v", err)
			return
		}
		wg.Done()
	}()
	wg.Wait()
}
