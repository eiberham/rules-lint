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
lint:
	go vet ./...
	go fmt ./...

.PHONY: deps
deps:
	go mod download
	go mod tidy

.PHONY: run
run: build
	./$(BUILD_DIR)/$(BINARY_NAME) ./rules

.PHONY: install-tools
install-tools:
	go install github.com/golang/mock/mockgen@latest

.PHONY: generate-mocks
generate-mocks:
	~/go/bin/mockgen -source=pkg/linter/types.go -destination=pkg/linter/mocks/mock_rule.go -package=mocks