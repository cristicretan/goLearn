
# Steps to Run Testing for the Kafka-based System

## 1. Start Kafka and Zookeeper

Ensure that Apache Kafka and Zookeeper are up and running. Use Docker Compose to start the necessary services.

**Command:**
```sh
docker-compose up -d
```

## 2. Start the System Consumers

Run the Go program that initializes and starts all the necessary consumers, such as the Trading Engine Consumer, Portfolio Manager Consumer, and Audit Logger.

**Command:**
```sh
go run main.go
```

## 3. Simulate Market Data and Orders

In separate terminal windows or tabs, run the following Python scripts to simulate market data and order commands:

### 3.1. Simulate Market Data

**Command:**
```sh
python simulate_market_data.py
```

### 3.2. Simulate Order Commands

**Command:**
```sh
python simulate_orders.py
```

## 4. Run the Test Suite

Execute the Go test suite to validate the system's functionality. This will run all the test cases defined in the test files.

**Command:**
```sh
go test -v
```

## 5. Check Test Results

After running the test suite, review the output for any passed or failed tests. The detailed output will help diagnose any issues or verify that the system operates as expected.
