.PHONY: help build run test clean migrate-up migrate-down docker-build docker-run

# Variables
APP_NAME=sitemapper
MAIN_PATH=./cmd/cli
BUILD_DIR=./bin
CONFIG_FILE=./configs/config.yaml

help: ## Show this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)
	@echo "Build complete: $(BUILD_DIR)/$(APP_NAME)"

run: ## Run the application
	@echo "Running $(APP_NAME)..."
	@go run $(MAIN_PATH)/main.go

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...

lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run ./...

tidy: ## Tidy go.mod
	@echo "Tidying go.mod..."
	@go mod tidy

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download

migrate-up: ## Run database migrations up
	@echo "Running migrations..."
	# TODO: Add migration command

migrate-down: ## Run database migrations down
	@echo "Rolling back migrations..."
	# TODO: Add migration rollback command

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t $(APP_NAME):latest .

docker-run: ## Run Docker container (interactive)
	@echo "Running Docker container..."
	@docker run -it --rm --env-file .env $(APP_NAME):latest

dev: ## Run in development mode with auto-reload
	@echo "Running in development mode..."
	@air || go run $(MAIN_PATH)/main.go

.DEFAULT_GOAL := help

