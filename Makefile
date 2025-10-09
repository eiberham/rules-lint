BINARY_NAME=rules-lint
BUILD_DIR=build
MAIN_PATH=./cmd/lint

.PHONY: help
help:
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

.PHONY: build-all
build-all:
	./scripts/build-all.sh

.PHONY: test
test:
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	go test -v -cover ./...

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: lint
lint: ## Run linter
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	@$$(go env GOPATH)/bin/golangci-lint run

.PHONY: dependencies
dependencies: ## Download dependencies and verify they are clean
	go mod download
	go mod tidy
	git diff --exit-code go.mod
	git diff --exit-code go.sum

.PHONY: run
run: build
	./$(BUILD_DIR)/$(BINARY_NAME) ./rules

.PHONY: install-tools
install-tools: ## Install development tools
	go install github.com/golang/mock/mockgen@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: generate-mocks
generate-mocks: ## Generate mocks
	~/go/bin/mockgen -source=pkg/linter/types.go -destination=pkg/linter/mock_rule.go -package=linter

.PHONY: ci
ci: deps lint test-coverage build ## Run all CI checks locally