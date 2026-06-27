BINARY_NAME=token-monitoring
DOCKER_IMAGE=midu16/token-monitoring:latest

.PHONY: help build test e2e-test deploy clean run

help: ## Display this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "; printf "\nUsage:  make \\n\n";while (FS = ":.*?## "; print $$2) printf "  %-20s %s\\n", $$1, $$2;} END {printf "\n"}'

build: ## Build the binary into the bin/ directory
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p bin
	@go build -o bin/$(BINARY_NAME) cmd/server/main.go

test: ## Run all unit tests
	@echo "Running unit tests..."
	@go test -v ./...

e2e-test: ## Run end-to-end tests (verifies the metrics endpoint is serving data)
	@echo "Starting E2E test environment..."
	@go run cmd/server/main.go & SERVER_PID=$$!; \
	sleep 5; \
	if curl -s http://localhost:8081/metrics | grep -q "llm_tokens_used_total"; then \
		echo "✅ E2E Test Passed: Metrics endpoint is reachable and contains data."; \
		kill $$SERVER_PID; \
	else \
		echo "❌ EMM E2E Test Failed: Metrics endpoint unreachable or empty."; \
		kill $$SERVER_PID; \
		exit 1; \
	fi

deploy: ## Build the Docker image for deployment
	@echo "Building Docker image $(DOCKER_IMAGE)..."
	@docker build -t $(DOCKER_IMAGE) .

run: ## Run the application locally
	@go run cmd/server/main.go

clean: ## Remove build artifacts and binaries
	@echo "Cleaning up..."
	@rm -rf bin
	@echo "Done."
