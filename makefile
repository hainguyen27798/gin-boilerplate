# Constants
APP_NAME := gin-boilerplate
BUILD_DIR := build
GO_SOURCE_FILES := $(shell find . -name '*.go' -not -path "./vendor/*")
GOLANGCI_LINT_CMD := golangci-lint run
GO_TEST_FLAGS := ./... -v

# Macros
CREATE_BUILD_DIR = @mkdir -p $(BUILD_DIR)
GO_RUN = air
GO_FMT = go fmt ./...
GO_CLEAN = rm -rf $(BUILD_DIR)

# Default target
.PHONY: all
all: build

# Run the application
.PHONY: run
run:
	$(GO_RUN)

# Build the application
.PHONY: build
build: clean $(GO_SOURCE_FILES)
	$(CREATE_BUILD_DIR)
	go build -ldflags="-X 'github.com/hainguyen27798/gin-boilerplate/version.Version=v1.0.1' \
                     -X 'github.com/hainguyen27798/gin-boilerplate/version.Commit=$(git rev-parse HEAD)' \
                     -X 'github.com/hainguyen27798/gin-boilerplate/version.BuildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ)'" -o $(BUILD_DIR)/$(APP_NAME) ./cmd

# Test the application
.PHONY: test
test:
	go test $(GO_TEST_FLAGS)

# Format the code
.PHONY: fmt
fmt:
	$(GO_FMT)

# Lint the code (requires golangci-lint)
.PHONY: lint
lint:
	$(GOLANGCI_LINT_CMD)

# Clean the build directory
.PHONY: clean
clean:
	$(GO_CLEAN)

# Install dependencies
.PHONY: deps
deps:
	go mod tidy
	go mod download

# Run everything
.PHONY: ci
ci: fmt lint test build