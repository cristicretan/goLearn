# Lab: Implementing a Pub/Sub System with Apache Kafka and Golang

## Objective
To implement a pub/sub system using Apache Kafka and Golang. This system will handle message queuing and pub/sub mechanisms, focusing on practical applications using Kafka. The lab will guide you through setting up Kafka, developing producers and consumers, and handling a real-world problem scenario.

## Lab Structure

### 1. Setting Up the Environment

#### 1.1 Kafka and Zookeeper Setup

**Description:** Apache Kafka requires Zookeeper for managing cluster coordination. We will use Docker to set up Kafka and Zookeeper quickly.

**Steps:**
1. **Create a `docker-compose.yml` file**: This file will define the services for Kafka and Zookeeper.
2. **Launch the services** using Docker Compose.

**Example `docker-compose.yml`:**
```yaml
version: '2'
services:
  zookeeper:
    image: wurstmeister/zookeeper:3.4.6
    ports:
      - "2181:2181"

  kafka:
    image: wurstmeister/kafka:latest
    ports:
      - "9092:9092"
    environment:
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
```

**Commands:**
- To start the services, run: `docker-compose up -d`
- To stop the services, run: `docker-compose down`

### 2. Implementing Producers

**File:** `producer.go`

**Description:** A producer sends messages to a Kafka topic. We'll use the Sarama library to interact with Kafka in Golang.

**Steps:**
1. **Install Sarama**: Use `go get github.com/Shopify/sarama` to install the Sarama library.
2. **Define the Producer**: Implement the `Producer` struct and methods to send messages to Kafka.

### 3. Implementing Consumers

**File:** `consumer.go`

**Description:** A consumer retrieves messages from a Kafka topic. We'll use the Sarama library for this as well.

**Steps:**
1. **Define the Consumer**: Implement the `Consumer` struct and methods to receive messages from Kafka.

## 4. Problem Statement: Implementing a Pub/Sub System for Order Notifications with Kafka

### Description
You are tasked with designing a pub/sub system using Kafka for an e-commerce platform's notification system. When a customer places an order, various subsystems need to be notified to take appropriate actions, such as updating inventory, processing payments, and preparing shipments.

### Requirements
- **Input**: An "Order Placed" message, containing order details such as order ID, customer ID, and list of items.
- **Output**: Notifications sent to multiple subsystems, ensuring each system receives the necessary data to perform its tasks.

### Inputs
1. "Order Placed: {order_id: 1001, customer_id: 501, items: [{id: 1, qty: 2}, {id: 2, qty: 1}]}"
2. "Order Placed: {order_id: 1002, customer_id: 502, items: [{id: 3, qty: 5}]}"

### Expected Outputs
1. **Inventory System**:
   - "Update Inventory: {order_id: 1001, items: [{id: 1, qty: -2}, {id: 2, qty: -1}]}"
   - "Update Inventory: {order_id: 1002, items: [{id: 3, qty: -5}]}"

2. **Billing System**:
   - "Process Payment: {order_id: 1001, amount: 150.00}"
   - "Process Payment: {order_id: 1002, amount: 200.00}"

3. **Shipping System**:
   - "Prepare Shipment: {order_id: 1001, customer_id: 501}"
   - "Prepare Shipment: {order_id: 1002, customer_id: 502}"

## Conclusion
This lab guides you through setting up and using Apache Kafka with Golang to implement a pub/sub system. You will learn to build producers and consumers, handle messages, and manage real-world scenarios like order notifications. Expand on this lab by adding features like error handling, message persistence, and using Kafka Streams for data processing.
