
# Homework Assignment: Advanced Pub/Sub System with Apache Kafka

## Title: **Developing an Advanced Pub/Sub System with Kafka**

## Objectives
To enhance understanding of Apache Kafka through implementing an advanced pub/sub system, focusing on message persistence, error handling, and scalability.

## Problem Statement
Build a notification system for a financial trading platform using Kafka. This system handles real-time market data, user actions (like placing orders), and sends notifications to various subsystems like trading engines, portfolio managers, and audit loggers.

## Scenario
1. **Market Data Publisher**: Publishes market data (e.g., stock prices).
2. **Order Processor**: Handles user orders and sends responses.
3. **Notification System**: Notifies subsystems in real-time.

## Tasks and Requirements

1. **Kafka Setup and Configuration**:
   - Multiple topics (`market-data`, `order-commands`, `notifications`).
   - Use multiple partitions for scalability.

2. **Producer Implementation**:
   - **Market Data Producer**: Publishes market data.
   - **Order Command Producer**: Sends user order commands.

3. **Consumer Implementation**:
   - **Trading Engine Consumer**: Processes orders.
   - **Portfolio Manager Consumer**: Updates portfolios.
   - **Audit Logger**: Logs transactions and market updates.

4. **Advanced Features**:
   - Message persistence and recovery.
   - Robust error handling.
   - Load balancing for consumers.

5. **Edge Cases**:
   - Non-existent stocks, duplicate orders, rate limiting.

6. **Testing and Validation**:
   - Comprehensive unit and integration tests for all scenarios.

## Skeletons

### market_data_producer.go
```go
package main

import (
    "github.com/Shopify/sarama"
    "log"
)

type MarketDataProducer struct {
    producer sarama.SyncProducer
}

func NewMarketDataProducer(brokers []string) (*MarketDataProducer, error) {
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true

    producer, err := sarama.NewSyncProducer(brokers, config)
    if err != nil {
        return nil, err
    }
    return &MarketDataProducer{producer: producer}, nil
}

func (p *MarketDataProducer) Publish(symbol string, price float64, timestamp int64) error {
    // TODO: Implement message publishing logic
    return nil
}

func main() {
    brokers := []string{"localhost:9092"}
    producer, err := NewMarketDataProducer(brokers)
    if err != nil {
        log.Fatalf("Failed to start producer: %v", err)
    }
    // TODO: Produce mock market data
}
```

### order_command_producer.go
```go
package main

import (
    "github.com/Shopify/sarama"
    "log"
)

type OrderCommandProducer struct {
    producer sarama.SyncProducer
}

func NewOrderCommandProducer(brokers []string) (*OrderCommandProducer, error) {
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true

    producer, err := sarama.NewSyncProducer(brokers, config)
    if err != nil {
        return nil, err
    }
    return &OrderCommandProducer{producer: producer}, nil
}

func (p *OrderCommandProducer) SendOrder(orderID, userID, symbol string, quantity int, orderType string, timestamp int64) error {
    // TODO: Implement order command publishing logic
    return nil
}

func main() {
    brokers := []string{"localhost:9092"}
    producer, err := NewOrderCommandProducer(brokers)
    if err != nil {
        log.Fatalf("Failed to start producer: %v", err)
    }
    // TODO: Produce mock order commands
}
```

### trading_engine_consumer.go
```go
package main

import (
    "github.com/Shopify/sarama"
    "log"
)

type TradingEngineConsumer struct {
    consumer sarama.Consumer
}

func NewTradingEngineConsumer(brokers []string) (*TradingEngineConsumer, error) {
    consumer, err := sarama.NewConsumer(brokers, nil)
    if err != nil {
        return nil, err
    }
    return &TradingEngineConsumer{consumer: consumer}, nil
}

func (c *TradingEngineConsumer) ProcessOrders() {
    // TODO: Implement order processing logic
}

func main() {
    brokers := []string{"localhost:9092"}
    consumer, err := NewTradingEngineConsumer(brokers)
    if err != nil {
        log.Fatalf("Failed to start consumer: %v", err)
    }
    consumer.ProcessOrders()
}
```

### portfolio_manager_consumer.go
```go
package main

import (
    "github.com/Shopify/sarama"
    "log"
)

type PortfolioManagerConsumer struct {
    consumer sarama.Consumer
}

func NewPortfolioManagerConsumer(brokers []string) (*PortfolioManagerConsumer, error) {
    consumer, err := sarama.NewConsumer(brokers, nil)
    if err != nil {
        return nil, err
    }
    return &PortfolioManagerConsumer{consumer: consumer}, nil
}

func (c *PortfolioManagerConsumer) UpdatePortfolio() {
    // TODO: Implement portfolio update logic
}

func main() {
    brokers := []string{"localhost:9092"}
    consumer, err := NewPortfolioManagerConsumer(brokers)
    if err != nil {
        log.Fatalf("Failed to start consumer: %v", err)
    }
    consumer.UpdatePortfolio()
}
```

### audit_logger.go
```go
package main

import (
    "github.com/Shopify/sarama"
    "log"
)

type AuditLogger struct {
    consumer sarama.Consumer
}

func NewAuditLogger(brokers []string) (*AuditLogger, error) {
    consumer, err := sarama.NewConsumer(brokers, nil)
    if err != nil {
        return nil, err
    }
    return &AuditLogger{consumer: consumer}, nil
}

func (c *AuditLogger) LogActivities() {
    // TODO: Implement logging of all messages
}

func main() {
    brokers := []string{"localhost:9092"}
    logger, err := NewAuditLogger(brokers)
    if err != nil {
        log.Fatalf("Failed to start logger: %v", err)
    }
    logger.LogActivities()
}
```

### kafka_setup.go
```go
package main

// TODO: Implement Kafka topic creation and configuration logic

func main() {
    // TODO: Setup Kafka topics and configurations
}
```

## Conclusion
This assignment challenges students to implement a sophisticated pub/sub system using Apache Kafka, emphasizing robust design, error handling, and testing. By completing this assignment, students will gain hands-on experience with Kafka's capabilities, preparing them for real-world applications in high-throughput, distributed systems.

## Submission Details
- **Deadline**: [Specify deadline]
- **Submission Format**: Submit a zip file containing all source files, a README.md with instructions, and a Makefile for building and testing the project.
- **Platform**: [Specify platform, e.g., university submission system]


## Test Cases

### kafka_pubsub_test.go

```go
package main

import (
    "testing"
)

func TestOrderPlacedNotification(t *testing.T) {
    brokers := []string{"localhost:9092"}
    producer, err := NewOrderCommandProducer(brokers)
    if err != nil {
        t.Fatalf("Failed to create producer: %v", err)
    }

    tradingConsumer, err := NewTradingEngineConsumer(brokers)
    if err != nil {
        t.Fatalf("Failed to create consumer: %v", err)
    }

    portfolioConsumer, err := NewPortfolioManagerConsumer(brokers)
    if err != nil {
        t.Fatalf("Failed to create consumer: %v", err)
    }

    auditLogger, err := NewAuditLogger(brokers)
    if err != nil {
        t.Fatalf("Failed to create logger: %v", err)
    }

    // Produce an "Order Placed" message
    err = producer.SendOrder("1001", "501", "AAPL", 10, "buy", 1625159076)
    if err != nil {
        t.Fatalf("Failed to send order: %v", err)
    }

    // Simulate consumers processing the message
    go tradingConsumer.ProcessOrders()
    go portfolioConsumer.UpdatePortfolio()
    go auditLogger.LogActivities()

    // Implement further checks to ensure correct processing and logging of messages
    // This typically involves capturing output from consumers and comparing it with expected results.
    // Note: Due to the asynchronous nature of Kafka, synchronization mechanisms or mocks may be required.
}

func TestDuplicateOrders(t *testing.T) {
    brokers := []string{"localhost:9092"}
    producer, err := NewOrderCommandProducer(brokers)
    if err != nil {
        t.Fatalf("Failed to create producer: %v", err)
    }

    // Send the same order twice
    orderID := "1002"
    err = producer.SendOrder(orderID, "502", "GOOGL", 15, "buy", 1625159077)
    if err != nil {
        t.Fatalf("Failed to send order: %v", err)
    }
    err = producer.SendOrder(orderID, "502", "GOOGL", 15, "buy", 1625159078)
    if err != nil {
        t.Fatalf("Failed to send duplicate order: %v", err)
    }

    // Verify that the system properly handles duplicate orders
    // This could involve checking logs, consumer outputs, or other validation mechanisms
}

func TestInvalidStockOrders(t *testing.T) {
    brokers := []string{"localhost:9092"}
    producer, err := NewOrderCommandProducer(brokers)
    if err != nil {
        t.Fatalf("Failed to create producer: %v", err)
    }

    // Send an order for a non-existent stock
    err = producer.SendOrder("1003", "503", "INVALID", 5, "buy", 1625159079)
    if err == nil {
        t.Fatalf("Expected error for invalid stock order, got nil")
    }

    // Verify the handling of invalid stock orders
    // This could involve checking error logs, consumer outputs, or other validation mechanisms
}

func TestRateLimiting(t *testing.T) {
    brokers := []string{"localhost:9092"}
    producer, err := NewMarketDataProducer(brokers)
    if err != nil {
        t.Fatalf("Failed to create producer: %v", err)
    }

    // Rapidly produce a large number of market data messages
    for i := 0; i < 1000; i++ {
        err = producer.Publish("AAPL", float64(i), 1625159080+int64(i))
        if err != nil {
            t.Fatalf("Failed to send market data: %v", err)
        }
    }

    // Verify that the system handles high throughput correctly and respects rate limiting
    // This could involve checking performance metrics, consumer outputs, or other validation mechanisms
}
```

## Makefile

```Makefile
# Variables
GO=go
BINARY_NAME=mykafka
SOURCE_DIR=./
TEST_DIR=./
TEST_OUTPUT=./test.out

# Default target
all: build

# Build the project
build:
    $(GO) build -o $(BINARY_NAME) $(SOURCE_DIR)

# Run the application
run: build
    ./$(BINARY_NAME)

# Clean the build files
clean:
    rm -f $(BINARY_NAME)
    rm -f $(TEST_OUTPUT)

# Run tests
test:
    $(GO) test $(TEST_DIR) -v | tee $(TEST_OUTPUT)

# Lint the code
lint:
    golangci-lint run

# Format the code
fmt:
    $(GO) fmt ./...

# Install necessary tools (for example, golangci-lint)
install-tools:
    $(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Phony targets
.PHONY: all build run clean test lint fmt install-tools
```

