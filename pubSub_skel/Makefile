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
	$(GO) test $(TEST_DIR)broker_test.go -v | tee $(TEST_OUTPUT)

test-produce-consume:
	$(GO) test $(TEST_DIR)broker_test.go -v -run ^TestProduceConsume$ | tee $(TEST_OUTPUT)

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
