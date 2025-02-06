# Constants
APP_NAME := gin-boilerplate
BUILD_DIR := build
GO_SOURCE_FILES := $(shell find . -name '*.go' -not -path "./vendor/*")
GOLANGCI_LINT_CMD := golangci-lint run
GO_TEST_FLAGS := ./... -v
GO_MODULE := $(shell grep "^module " go.mod | sed -E 's/module (.*)/\1/')
GO_VERSION_FILE := ./metadata/version.txt

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
build: clean $(GO_SOURCE_FILES) increment-version
	$(eval VERSION := $(shell cat $(GO_VERSION_FILE))) \
	echo "Building version $(VERSION)..."
	$(CREATE_BUILD_DIR)
	go build \
		-ldflags="-X '$(GO_MODULE)/metadata.Version=$(VERSION)' \
              -X '$(GO_MODULE)/metadata.Commit=$(shell git rev-parse --short=8 HEAD)' \
            	-X '$(GO_MODULE)/metadata.BuildDate=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)'" \
  	-o $(BUILD_DIR)/$(APP_NAME) ./cmd

# Test the application
.PHONY: test
test:
	go clean -testcache && \
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

.PHONY: increment-version
increment-version:
	@if [ ! -f $(GO_VERSION_FILE) ]; then \
		touch $(GO_VERSION_FILE); \
		chmod 644 $(GO_VERSION_FILE); \
		echo "v0.0.0" > $(GO_VERSION_FILE); \
	fi
	@CURRENT_VERSION=$$(cat $(GO_VERSION_FILE) | sed 's/v//'); \
	MAJOR=$$(echo "$$CURRENT_VERSION" | awk -F. '{print $$1}'); \
	MINOR=$$(echo "$$CURRENT_VERSION" | awk -F. '{print $$2}'); \
	PATCH=$$(echo "$$CURRENT_VERSION" | awk -F. '{print $$3}'); \
	BUMP=$(BUMP); \
	if [ -z "$$BUMP" ]; then BUMP="patch"; fi; \
	case $$BUMP in \
		major) \
			NEW_MAJOR=$$((MAJOR + 1)); \
			NEW_VERSION="v$${NEW_MAJOR}.0.0";; \
		minor) \
			NEW_MINOR=$$((MINOR + 1)); \
			NEW_VERSION="v$${MAJOR}.$${NEW_MINOR}.0";; \
		patch) \
			NEW_PATCH=$$((PATCH + 1)); \
			NEW_VERSION="v$${MAJOR}.$${MINOR}.$${NEW_PATCH}";; \
		*) echo "Invalid bump type. Use major/minor/patch"; exit 1;; \
	esac; \
	echo "$$NEW_VERSION" > $(GO_VERSION_FILE)
