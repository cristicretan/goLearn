package main

import (
	"fmt"
	"sync"
)

type Message struct {
	ID      int
	Content string
	Acked   bool
}

// MyKafka represents the broker structure
type MyKafka struct {
	mu        sync.Mutex
	messages  []Message
	nextMsgID int
	consumers map[string]bool // list of consumers
	//currentConsumer int      // index of the next consumer to receive a message
}

var myKafka *MyKafka
var once sync.Once

func GetBrokerInstance() *MyKafka {
	once.Do(func() {
		myKafka = NewMyKafka()
	})
	return myKafka
}

// NewMyKafka initializes a new broker instance
func NewMyKafka() *MyKafka {
	return &MyKafka{
		messages:  []Message{},
		mu:        sync.Mutex{},
		nextMsgID: 0,
		consumers: make(map[string]bool),
	}
}

// Produce adds a message to the broker's queue
func (b *MyKafka) Produce(content string) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if content == "" {
		return fmt.Errorf("message cannot be empty")
	}

	msg := Message{
		ID:      b.nextMsgID,
		Content: content,
		Acked:   false,
	}

	b.messages = append(b.messages, msg)
	b.nextMsgID++
	return nil
}

// Consume retrieves and removes the first message from the queue
func (b *MyKafka) Consume(consumerID string) (Message, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.messages) == 0 {
		return Message{}, fmt.Errorf("no messages to consume")
	}
	err := b.Acknowledge(consumerID, b.messages[0].ID)
	fmt.Printf("Message %d consumed by %s\n", b.messages[0].ID, consumerID)
	if err != nil {
		return Message{}, fmt.Errorf("error acknowledging message. err: %v", err)
	}
	ans := b.messages[0]
	b.messages = b.messages[1:]
	return ans, nil
}

// Acknowledge marks a message as processed
func (b *MyKafka) Acknowledge(consumerID string, msgID int) error {
	for i := range b.messages {
		if b.messages[i].ID == msgID {
			b.messages[i].Acked = true
			return nil
		}
	}
	return fmt.Errorf("message ID %d not found", msgID)
}

// RegisterConsumer registers a new consumer with the broker
func (b *MyKafka) RegisterConsumer(consumerID string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.consumers[consumerID]; !exists {
		b.consumers[consumerID] = true
		fmt.Printf("consumer %s registered\n", consumerID)
	}
}

// LoadBalancing method to distribute messages across consumers
func (b *MyKafka) LoadBalancing() {
	// TODO: Implement simple round-robin or another load balancing strategy
}

func main() {
	broker := GetBrokerInstance()
	err := broker.Produce("blabla")
	if err != nil {
		panic("Could not produce message")
	}

	consumerID := "consumer1"
	broker.RegisterConsumer(consumerID)
	msg, err := broker.Consume(consumerID)
	if err != nil {
		panic("Could not consume message")
	}
	fmt.Printf("Consumer %s consumed %s\n", consumerID, msg.Content)
	consumerID2 := "consumer2"
	broker.RegisterConsumer(consumerID2)
	msg, err = broker.Consume(consumerID)
	if err != nil {
		fmt.Printf("error consuming message. err: %v", err)
	}
	fmt.Printf("Consumer %s consumed %s\n", consumerID, msg.Content)
}
