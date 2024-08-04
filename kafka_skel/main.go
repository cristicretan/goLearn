package main

import (
	"fmt"
	"log"
	"time"
)

const (
	invetorySystem = "invetory-system"
	billingSystem  = "billing-system"
	shippingSystem = "shipping-system"
)

type Item struct {
	ID  int
	qty int
}

type Order struct {
	orderID    int
	customerID int
	items      []Item
}

func (o Order) String() string {
	itemsStr := "["
	for i := 0; i < len(o.items)-1; i++ {
		itemsStr += fmt.Sprintf("{id: %d, qty: %d}, ", o.items[i].ID, o.items[i].qty)
	}
	itemsStr += fmt.Sprintf("{id: %d, qty: %d}]", o.items[len(o.items)-1].ID, o.items[len(o.items)-1].qty)
	return fmt.Sprintf("{orderID: %d, customerID: %d, items: %s}", o.orderID, o.customerID, itemsStr)
}

var brokers = []string{"localhost:9092"}
var topic = "test-topic"

func placeOrder(order Order) {

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
