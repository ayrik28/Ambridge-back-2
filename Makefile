.PHONY: run build clean

# Default target
all: run

# Run the server
run:
	go run cmd/server/main.go

# Build the server
build:
	go build -o bin/server cmd/server/main.go

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod download

# Create .env file from sample if it doesn't exist
env:
	@if [ ! -f .env ]; then \
		cp env.sample .env; \
		echo ".env file created from env.sample"; \
	else \
		echo ".env file already exists"; \
	fi
