/*
* @Author: Cristi Cretan
* @Date:   2024-08-03 19:19:59
* @Last Modified by:   Cristi Cretan
* @Last Modified time: 2024-08-03 19:20:10
*/
package main

import (
    "testing"
)

func TestProduceConsume(t *testing.T) {
    broker := NewMyKafka()

    // Test producing a message
    broker.Produce("Order Placed: {order_id: 123, customer_id: 456, items: [{id: 1, qty: 2}, {id: 2, qty: 1}]}")
    if len(broker.messages) != 1 {
        t.Errorf("Expected 1 message, got %d", len(broker.messages))
    }

    // Test consuming a message
    msg := broker.Consume()
    if msg != "Order Placed: {order_id: 123, customer_id: 456, items: [{id: 1, qty: 2}, {id: 2, qty: 1}]}" {
        t.Errorf("Expected message 'Order Placed', got %s", msg)
    }

    // Test empty queue handling
    msg = broker.Consume()
    if msg != "" {
        t.Errorf("Expected empty string, got %s", msg)
    }
}

func TestAcknowledge(t *testing.T) {
    broker := NewMyKafka()
    broker.Produce("Order Placed: {order_id: 124, customer_id: 457, items: [{id: 3, qty: 4}]}")
    broker.Acknowledge("Order Placed: {order_id: 124, customer_id: 457, items: [{id: 3, qty: 4}]}")
    // Verify that the message is acknowledged properly
}

func TestLoadBalancing(t *testing.T) {
    broker := NewMyKafka()
    broker.RegisterConsumer("Consumer1")
    broker.RegisterConsumer("Consumer2")
    broker.Produce("Order Placed: {order_id: 125, customer_id: 458, items: [{id: 4, qty: 5}]}")

    // Implement logic to check if the message is distributed correctly
}
