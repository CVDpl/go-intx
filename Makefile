# Simple Makefile for go-intx

.PHONY: all build test bench clean dev help

# Default target: build and test
all: build test

# Build the project
build:
	@echo "Building..."
	@go build ./...

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@go clean
	@rm -f *.out *.prof

# Development: build and test with race detector
dev:
	@echo "Building and testing with race detector..."
	@go build -race ./...
	@go test -race -v ./...

# Show help
help:
	@echo "Available targets:"
	@echo "  all     - Build and test the project"
	@echo "  build   - Build the project"
	@echo "  test    - Run tests"
	@echo "  bench   - Run benchmarks"
	@echo "  clean   - Clean build artifacts"
	@echo "  dev     - Build and test with race detector"
	@echo "  help    - Show this help message" 