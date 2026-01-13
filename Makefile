.PHONY: build test clean install docker-build docker-multiarch help

# Build variables
BINARY_NAME=drone-ai-review
DOCKER_IMAGE=drone-ai-review
DOCKER_REGISTRY=your-registry

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the binary
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) .
	@echo "Build complete: $(BINARY_NAME)"

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -cover ./plugin/
	go test -coverprofile=coverage.out ./plugin/
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html
	rm -rf test-output/
	@echo "Clean complete"

install: build ## Install the binary
	@echo "Installing $(BINARY_NAME)..."
	install -m 755 $(BINARY_NAME) /usr/local/bin/
	@echo "Install complete"

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE):latest .
	@echo "Docker build complete"

docker-multiarch: ## Build multi-architecture Docker images
	@echo "Building multi-architecture Docker images..."
	./build-multiarch.sh
	@echo "Multi-arch build complete"

docker-push: ## Build and push multi-architecture images to registry
	@echo "Building and pushing to registry..."
	PUSH=true REGISTRY=$(DOCKER_REGISTRY) IMAGE_NAME=$(DOCKER_IMAGE) ./build-multiarch.sh
	@echo "Push complete"

run-tests: test ## Alias for test

fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...
	@echo "Format complete"

lint: ## Run linter (requires golangci-lint)
	@echo "Running linter..."
	golangci-lint run ./...
	@echo "Lint complete"

mod-tidy: ## Tidy go modules
	@echo "Tidying modules..."
	go mod tidy
	@echo "Modules tidied"

all: clean fmt mod-tidy test build ## Run all: clean, format, tidy, test, and build

