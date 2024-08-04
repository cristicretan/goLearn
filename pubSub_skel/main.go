package main

import (
	"errors"
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
	mu              sync.Mutex
	messages        []Message
	nextMsgID       int
	consumers       []string // list of consumers
	consumers_set   map[string]bool
	currentConsumer int // index of the next consumer to receive a message
}

var myKafka *MyKafka
var once sync.Once

type ConsumeOptions struct {
	ConsumerID string
}

type Option func(*ConsumeOptions)

func WithConsumerID(consumerID string) Option {
	return func(o *ConsumeOptions) {
		o.ConsumerID = consumerID
	}
}

func GetBrokerInstance() *MyKafka {
	once.Do(func() {
		myKafka = NewMyKafka()
	})
	return myKafka
}

// NewMyKafka initializes a new broker instance
func NewMyKafka() *MyKafka {
	return &MyKafka{
		messages:      []Message{},
		mu:            sync.Mutex{},
		nextMsgID:     0,
		consumers:     []string{},
		consumers_set: map[string]bool{},
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
func (b *MyKafka) Consume(opts ...Option) (Message, error) {
	var options ConsumeOptions
	for _, opt := range opts {
		opt(&options)
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.messages) == 0 {
		return Message{}, fmt.Errorf("no messages to consume")
	}

	if options.ConsumerID != "" {
		err := b.Acknowledge(options.ConsumerID, b.messages[0].ID)
		if _, exists := b.consumers_set[options.ConsumerID]; !exists {
			return Message{}, fmt.Errorf("consumer %s does not exist", options.ConsumerID)
		}
		if err != nil {
			return Message{}, fmt.Errorf("error acknowledging message. err: %v", err)
		}
		ans := b.messages[0]
		b.messages = b.messages[1:]
		return ans, nil
	}

	if len(b.consumers) == 0 {
		return Message{}, fmt.Errorf("no consumers registered")
	}
	consumerID := b.consumers[b.currentConsumer]
	fmt.Printf("Assigning message to consumer: %s\n", consumerID)

	b.currentConsumer = (b.currentConsumer + 1) % len(b.consumers)
	for _, msg := range b.messages {
		if !msg.Acked {
			err := b.Acknowledge(consumerID, msg.ID)
			if err != nil {
				return Message{}, fmt.Errorf("error acknowledging message. err: %v", err)
			}
			b.messages = b.messages[1:]
			return msg, nil
		}
	}
	return Message{}, errors.New("no unacknowledged messages available")
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
	if _, exists := b.consumers_set[consumerID]; !exists {
		b.consumers_set[consumerID] = true
		b.consumers = append(b.consumers, consumerID)
		fmt.Printf("Consumer %s registered\n", consumerID)
	}
}

//
//// LoadBalancing method to distribute messages across consumers
//func (b *MyKafka) LoadBalancing() {
//	// TODO: Implement simple round-robin or another load balancing strategy
//
//}

func main() {
	broker := GetBrokerInstance()
	err := broker.Produce("1 message")
	if err != nil {
		panic("Could not produce message")
	}
	err = broker.Produce("2 message")
	if err != nil {
		panic("Could not produce message")
	}
	err = broker.Produce("3 message")
	if err != nil {
		panic("Could not produce message")
	}
	err = broker.Produce("4 message")
	if err != nil {
		panic("Could not produce message")
	}
	err = broker.Produce("5 message")
	if err != nil {
		panic("Could not produce message")
	}

	consumerID := "consumer1"
	broker.RegisterConsumer(consumerID)
	msg, err := broker.Consume(WithConsumerID(consumerID))
	if err != nil {
		panic("Could not consume message")
	}
	fmt.Printf("Consumer %s consumed %s\n", consumerID, msg.Content)
	consumerID2 := "consumer2"
	broker.RegisterConsumer(consumerID2)
	msg, err = broker.Consume(WithConsumerID(consumerID))
	if err != nil {
		fmt.Printf("error consuming message. err: %v", err)
	}

	msg, err = broker.Consume()
	if err != nil {
		fmt.Printf("error consuming message. err: %v", err)
	}

	msg, err = broker.Consume()
	if err != nil {
		fmt.Printf("error consuming message. err: %v", err)
	}

	msg, err = broker.Consume()
	if err != nil {
		fmt.Printf("error consuming message. err: %v", err)
	}
}
