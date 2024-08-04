package main

// Consumer represents a message consumer
type Consumer struct {
	ID string
}

// NewConsumer initializes a new consumer instance
func NewConsumer(id string) *Consumer {
	broker := GetBrokerInstance()
	broker.RegisterConsumer(id)
	return &Consumer{ID: id}
}

// Receive retrieves a message from the broker
func (c *Consumer) Receive() (string, error) {
	myKafka := GetBrokerInstance()
	consume, err := myKafka.Consume(WithConsumerID(c.ID))
	if err != nil {
		return "", err
	}
	return consume.Content, nil
}
