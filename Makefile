# Makefile for ROR (Release Operate Report) project

.PHONY: help setup build test clean run dev docs docker chart lint format deps update-deps

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

# Variables
GO_VERSION := $(shell go version | cut -d' ' -f3)
PROJECT_NAME := ror
DOCKER_COMPOSE := docker compose
HELM := helm
KUBECTL := kubectl

# Check if required tools are installed
check-tools:
	@command -v go >/dev/null 2>&1 || { echo "Go is required but not installed. Aborting." >&2; exit 1; }
	@command -v docker >/dev/null 2>&1 || { echo "Docker is required but not installed. Aborting." >&2; exit 1; }
	@command -v $(DOCKER_COMPOSE) >/dev/null 2>&1 || { echo "Docker Compose is required but not installed. Aborting." >&2; exit 1; }

##@ Development Commands

setup: check-tools ## Initial project setup and install dependencies
	@echo "Setting up ROR development environment..."
	@echo "Go version: $(GO_VERSION)"
	@go mod download
	@go mod verify
	@./r.sh setup 2>/dev/null || echo "Running initial setup script..."
	@echo "Setup complete! Run 'make run' to start development infrastructure."

build: ## Build all Go binaries
	@echo "Building Go binaries..."
	@go build -v ./cmd/...
	@echo "Build complete!"

test: ## Run all tests
	@echo "Running tests..."
	@go test -v ./...
	@echo "Tests complete!"

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

clean: ## Clean build artifacts and containers
	@echo "Cleaning build artifacts..."
	@go clean -cache
	@go clean -modcache
	@rm -f coverage.out coverage.html
	@$(DOCKER_COMPOSE) down -v 2>/dev/null || true
	@docker system prune -f 2>/dev/null || true
	@echo "Clean complete!"

run: check-tools ## Start development infrastructure
	@echo "Starting ROR development infrastructure..."
	@./r.sh

dev: check-tools ## Start development with additional services
	@echo "Starting ROR development with additional services..."
	@./r.sh jaeger opentelemetry-collector

##@ Infrastructure Management

infra-up: check-tools ## Start core infrastructure (MongoDB, Vault, etc.)
	@echo "Starting core infrastructure..."
	@$(DOCKER_COMPOSE) --env-file .env up -d openldap dex init-dex-db vault mongodb rabbitmq mongo-express valkey

infra-down: ## Stop infrastructure
	@echo "Stopping infrastructure..."
	@$(DOCKER_COMPOSE) down

infra-reset: ## Reset infrastructure (clean volumes)
	@echo "Resetting infrastructure..."
	@$(DOCKER_COMPOSE) down -v
	@docker volume prune -f

##@ Code Quality

format: ## Format Go code
	@echo "Formatting Go code..."
	@go fmt ./...
	@echo "Code formatted!"

lint: ## Run Go linters
	@echo "Running Go linters..."
	@command -v golangci-lint >/dev/null 2>&1 || { echo "Installing golangci-lint..."; go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; }
	@golangci-lint run ./...
	@echo "Linting complete!"

deps: ## Download and verify dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod verify
	@go mod tidy
	@echo "Dependencies updated!"

update-deps: ## Update dependencies
	@echo "Updating dependencies..."
	@go get -u ./...
	@go mod tidy
	@echo "Dependencies updated!"

##@ Documentation

docs-build: ## Build documentation
	@echo "Building documentation..."
	@command -v mkdocs >/dev/null 2>&1 || { echo "MkDocs is required. Install with: pip install -r mkdocs-requirements.txt"; exit 1; }
	@./cmd/docs/collectdocs.sh
	@mkdocs build
	@echo "Documentation built in site/ directory"

docs-serve: ## Serve documentation locally
	@echo "Serving documentation..."
	@command -v mkdocs >/dev/null 2>&1 || { echo "MkDocs is required. Install with: pip install -r mkdocs-requirements.txt"; exit 1; }
	@./cmd/docs/collectdocs.sh
	@mkdocs serve

docs: docs-build docs-serve ## Build and serve documentation

##@ Docker & Kubernetes

docker-build: ## Build Docker images
	@echo "Building Docker images..."
	@docker build -t $(PROJECT_NAME):latest .
	@echo "Docker images built!"

chart-lint: ## Lint Helm charts
	@echo "Linting Helm charts..."
	@for chart in charts/*/; do \
		if [ -f "$$chart/Chart.yaml" ]; then \
			echo "Linting $$chart"; \
			$(HELM) lint "$$chart"; \
		fi \
	done
	@echo "Chart linting complete!"

chart-test: ## Test Helm charts
	@echo "Testing Helm charts..."
	@for chart in charts/*/; do \
		if [ -f "$$chart/Chart.yaml" ]; then \
			echo "Testing $$chart"; \
			$(HELM) template test "$$chart" > /dev/null; \
		fi \
	done
	@echo "Chart testing complete!"

##@ Development Utilities

generate-ts: ## Generate TypeScript models
	@echo "Generating TypeScript models..."
	@if [ -d "typescript/models" ]; then \
		go run cmd/generator/main.go; \
	fi
	@echo "TypeScript models generated!"

logs: ## Follow Docker Compose logs
	@echo "Following Docker Compose logs..."
	@$(DOCKER_COMPOSE) logs -f

##@ Security & Performance

security-scan: ## Run security scan
	@echo "Running security scan..."
	@command -v gosec >/dev/null 2>&1 || { echo "Installing gosec..."; go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest; }
	@gosec ./...
	@echo "Security scan complete!"

bench: ## Run benchmarks
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./...
	@echo "Benchmarks complete!"

install-tools: ## Install development tools
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
	@echo "Development tools installed!"

info: ## Show project information
	@echo "ROR (Release Operate Report) Project Information:"
	@echo "================================================"
	@echo "Go Version: $(GO_VERSION)"
	@echo "Project: $(PROJECT_NAME)"
	@echo "Repository: github.com/NorskHelsenett/ror"
	@echo ""
	@echo "Available charts:"
	@find charts -name "Chart.yaml" -exec dirname {} \; | sed 's/charts\//  - /'
	@echo ""
	@echo "Services:"
	@echo "  - ROR API"
	@echo "  - ROR Web"
	@echo "  - Multiple microservices (auth, nhn, slack, tanzu, etc.)"
	@echo "  - Infrastructure (MongoDB, Vault, RabbitMQ, etc.)"