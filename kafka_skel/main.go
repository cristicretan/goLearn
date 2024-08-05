package main

import (
	"fmt"
	"log"
	"time"
)

const (
	inventorySystemTopic = "inventory-system"
	billingSystemTopic   = "billing-system"
	shippingSystemTopic  = "shipping-system"
)

type Item struct {
	ID  int
	Qty int
}

type Order struct {
	OrderID    int
	CustomerID int
	Items      []Item
}

func (o Order) String() string {
	itemsStr := "["
	for i := 0; i < len(o.Items)-1; i++ {
		itemsStr += fmt.Sprintf("{id: %d, qty: %d}, ", o.Items[i].ID, o.Items[i].Qty)
	}
	itemsStr += fmt.Sprintf("{id: %d, qty: %d}]", o.Items[len(o.Items)-1].ID, o.Items[len(o.Items)-1].Qty)
	return fmt.Sprintf("{orderID: %d, customerID: %d, items: %s}", o.OrderID, o.CustomerID, itemsStr)
}

var brokers = []string{"localhost:9092"}
var topic = "test-topic"

type InventorySystem struct {
	Items    []Item
	Consumer *Consumer
}

func newInventorySystem() *InventorySystem {
	consumer, err := NewConsumer(brokers)
	if err != nil {
		log.Fatalf("Error creating invetory system: %v", err)
	}
	return &InventorySystem{
		Items:    []Item{},
		Consumer: consumer,
	}
}

func (s *InventorySystem) Start() {
	go func() {
		for {
			err := s.Consumer.Receive(inventorySystemTopic)
			if err != nil {
				log.Printf("Error consuming message in invetory system: %v", err)
			}
		}
	}()
}

type BillingSystem struct {
	Balance  float64
	Consumer *Consumer
}

func newBillingSystem(balance float64) *BillingSystem {
	consumer, err := NewConsumer(brokers)
	if err != nil {
		log.Fatalf("Error creating billing system: %v", err)
	}
	return &BillingSystem{
		Balance:  balance,
		Consumer: consumer,
	}
}

func (s *BillingSystem) Start() {
	go func() {
		for {
			err := s.Consumer.Receive(billingSystemTopic)
			if err != nil {
				log.Printf("Error consuming message in billing system: %v", err)
			}
		}
	}()
}

type Customer struct {
	ID      int
	Address string
}

type ShippingSystem struct {
	Customers []Customer
	Consumer  *Consumer
}

func newShippingSystem() *ShippingSystem {
	consumer, err := NewConsumer(brokers)
	if err != nil {
		log.Fatalf("Error creating shipping system: %v", err)
	}
	return &ShippingSystem{
		Customers: []Customer{},
		Consumer:  consumer,
	}
}

func (s *ShippingSystem) Start() {
	go func() {
		for {
			err := s.Consumer.Receive(shippingSystemTopic)
			if err != nil {
				log.Printf("Error consuming message in shipping system: %v", err)
			}
		}
	}()
}

func placeOrder(producer *Producer, order Order) error {
	err := producer.Send(inventorySystemTopic, order.String())
	if err != nil {
		log.Fatalf("Error placing order in invetory system: %v", err)
		return err
	}
	err = producer.Send(billingSystemTopic, order.String())
	if err != nil {
		log.Fatalf("Error placing order in billing system: %v", err)
		return err
	}
	err = producer.Send(shippingSystemTopic, order.String())
	if err != nil {
		log.Fatalf("Error placing order in shipping system: %v", err)
		return err
	}
	return nil
}

func main() {
	inventorySystem := newInventorySystem()
	billingSystem := newBillingSystem(1000)
	shippingSystem := newShippingSystem()

	inventorySystem.Start()
	billingSystem.Start()
	shippingSystem.Start()

	time.Sleep(time.Second * 10)
}
