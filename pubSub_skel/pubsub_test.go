/*
* @Author: Cristi Cretan
* @Date:   2024-08-03 19:20:20
* @Last Modified by:   Cristi Cretan
* @Last Modified time: 2024-08-03 19:20:30
 */
package main

//
//func TestOrderPlacedNotification(t *testing.T) {
//	broker := NewMyKafka()
//	// Register consumers for different subsystems
//	broker.RegisterConsumer("InventorySystem")
//	broker.RegisterConsumer("BillingSystem")
//	broker.RegisterConsumer("ShippingSystem")
//
//	// Produce an "Order Placed" message
//	broker.Produce("Order Placed: {order_id: 1001, customer_id: 501, items: [{id: 1, qty: 2}, {id: 2, qty: 1}]}")
//
//	// Verify that the appropriate notifications are sent to each consumer
//	// Inventory System check
//	invMsg, err := broker.Consume()
//	if err != nil {
//		t.Fatalf("Failed to consume message: %s", err)
//	}
//	expectedInvMsg := "Update Inventory: {order_id: 1001, items: [{id: 1, qty: -2}, {id: 2, qty: -1}]}"
//	if invMsg != expectedInvMsg {
//		t.Errorf("Expected message %s, got %s", expectedInvMsg, invMsg)
//	}
//
//	// Billing System check
//	billingMsg, err := broker.Consume()
//	if err != nil {
//		t.Fatalf("Failed to consume message: %s", err)
//	}
//	expectedBillingMsg := "Process Payment: {order_id: 1001, amount: 150.00}"
//	if billingMsg != expectedBillingMsg {
//		t.Errorf("Expected message %s, got %s", expectedBillingMsg, billingMsg)
//	}
//
//	// Shipping System check
//	shipMsg, err := broker.Consume()
//	if err != nil {
//		t.Fatalf("Failed to consume message: %s", err)
//	}
//	expectedShipMsg := "Prepare Shipment: {order_id: 1001, customer_id: 501}"
//	if shipMsg != expectedShipMsg {
//		t.Errorf("Expected message %s, got %s", expectedShipMsg, shipMsg)
//	}
//}
