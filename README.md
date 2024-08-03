# Lab: Building "MyKafka" - A Simplified Message Broker System with Golang

## Objective
Create a simplified message broker system called "MyKafka" using Golang. This system will handle basic message queuing and pub/sub mechanisms. We will implement a broker, producer, and consumer, and test them using a real-world problem scenario.

## Lab Structure

### 1. System Architecture

- **Components**: 
  - **MyKafka Broker**: Handles message storage and delivery.
  - **Producers**: Send messages to the broker.
  - **Consumers**: Retrieve messages from the broker.

### 2. Implementing MyKafka Broker

**File:** `broker.go`

**Description:** The `MyKafka` broker is the central component of our message broker system. It will manage the storage and retrieval of messages. We will use an in-memory data structure for simplicity, but the design will allow for easy extension to use persistent storage (like BoltDB or Badger) if needed.

**Steps:**
1. **Define the Broker Structure**: The `MyKafka` struct will include a mutex for thread safety and a slice to store messages.
2. **Initialize Broker**: Implement the `NewMyKafka` function to initialize the broker.
3. **Message Handling**:
   - `Produce`: Add a message to the broker's queue. Use the mutex to ensure thread safety.
   - `Consume`: Retrieve and remove a message from the queue, ensuring proper synchronization with the mutex.
4. **Additional Features**:
   - **Acknowledgment**: Implement an `Acknowledge` method to mark messages as processed. This can be used to manage the state of messages and ensure reliability.
   - **Consumer Management**: Allow consumers to register with the broker and implement load balancing strategies, such as round-robin.

### 3. Implementing Producers

**File:** `producer.go`

**Description:** The producer sends messages to the broker. This component will connect to the `MyKafka` broker and produce messages, which could represent various data, such as logs, user actions, or system events.


### 4. Implementing Consumers

**File:** `consumer.go`

**Description:** The consumer retrieves messages from the broker. This component will connect to the `MyKafka` broker and consume messages, processing them according to the application logic.

**Steps:**
1. **Define the Consumer Structure**: Include necessary fields for broker connection details.
2. **Initialize Consumer**: Implement the `NewConsumer` function to set up the consumer with the broker's address.
3. **Receiving Messages**: Implement the `Receive` method to consume messages from the broker. Handle the message processing logic, such as updating a database, triggering actions, or logging.

## 5. Testing and Validation

### 5.1. Real-World Problem Scenario

**Scenario:** An online retailer uses MyKafka to handle order processing. When a customer places an order, the producer sends an "Order Placed" message. A consumer reads this message, processes the order (e.g., updating inventory, notifying shipping), and sends an "Order Confirmed" message.

**Messages:**
1. "Order Placed: {order_id: 123, customer_id: 456, items: [{id: 1, qty: 2}, {id: 2, qty: 1}]}"
2. "Order Confirmed: {order_id: 123, status: confirmed, expected_delivery: '2024-08-10'}"

**Test Cases:**

- **TestProduceConsume:** This test case will produce a message to the broker and then consume it, verifying that the messages are correctly stored and retrieved.
- **TestAcknowledge:** This test case will simulate acknowledging a message, ensuring that the broker handles the message state appropriately.
- **TestLoadBalancing:** This test case will register multiple consumers and test the broker's load balancing strategy.

## Conclusion
This lab provides the foundation to build a simplified message broker system called "MyKafka" with Golang. By following the skeletons and implementing the code, you will learn how to manage message queues, produce and consume messages, and handle real-world scenarios such as order processing in an e-commerce system. Consider extending this system with features like message persistence, advanced consumer group management, and support for different serialization formats.

## 6. Problem Statement: Implementing a Pub/Sub System for Order Notifications

### Description
You are tasked with designing a pub/sub system to handle notifications for an e-commerce platform. When a customer places an order, various subsystems need to be notified to take appropriate actions. For example, the inventory system needs to update stock levels, the billing system needs to process payments, and the shipping system needs to prepare for dispatch.

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

