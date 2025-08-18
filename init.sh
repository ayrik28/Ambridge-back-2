#!/bin/bash

echo "Initializing Ambridge Backend project..."

# Download dependencies
echo "Downloading dependencies..."
go mod tidy

# Create .env file if it doesn't exist
if [ ! -f .env ]; then
    echo "Creating .env file from env.sample..."
    cp env.sample .env
    echo ".env file created. Please update it with your configuration."
else
    echo ".env file already exists."
fi

# Create .gitignore if it doesn't exist
if [ ! -f .gitignore ]; then
    echo "Creating .gitignore file from .gitignore.sample..."
    cp .gitignore.sample .gitignore
    echo ".gitignore file created."
else
    echo ".gitignore file already exists."
fi

echo "Initialization completed successfully!"
echo "You can now run the server with: go run cmd/server/main.go"
echo "Or use the Makefile: make run"
