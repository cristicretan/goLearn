package main

import (
	"fmt"
	"log"
	"time"
)

const (
	invetorySystemTopic = "invetory-system"
	billingSystemTopic  = "billing-system"
	shippingSystemTopic = "shipping-system"
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

func placeOrder(producer *Producer, order Order) error {
	err := producer.Send(invetorySystemTopic, order.String())
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
	producer, err := NewProducer(brokers)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}

	err = producer.Send(topic, "Hello, Kafka!")
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	log.Printf("Message sent to topic %s", topic)

	consumer, err := NewConsumer(brokers)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}

	go func() {
		err = consumer.Receive(topic)
		if err != nil {
			log.Fatalf("Failed to receive message: %v", err)
		}
	}()

	time.Sleep(time.Second * 10)
}
