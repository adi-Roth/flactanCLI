# Project settings
PROJECT_NAME := flactanCLI
SRC_DIR := .
BIN_DIR := bin
TEST_DIR := tests

# Go settings
GO := go
GO_BUILD := $(GO) build
GO_TEST := $(GO) test
GO_CLEAN := $(GO) clean

# Build targets
all: build

build:
	@echo "Building $(PROJECT_NAME)..."
	@mkdir -p $(BIN_DIR)
	@$(GO_BUILD) -o $(BIN_DIR)/$(PROJECT_NAME) $(SRC_DIR)/main.go
	@echo "Build complete!"

build-windows:
	@echo "Building Windows binary..."
	@mkdir -p $(BIN_DIR)
	GOOS=windows GOARCH=amd64 $(GO_BUILD) -o $(BIN_DIR)/$(PROJECT_NAME).exe $(SRC_DIR)/main.go
	@echo "Windows build complete!"

build-linux:
	@echo "Building Linux binary..."
	@mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(BIN_DIR)/$(PROJECT_NAME) $(SRC_DIR)/main.go
	@echo "Linux build complete!"

test:
	@echo "Running tests..."
	@$(GO_TEST) -v ./$(TEST_DIR)/...
	@echo "All tests passed!"

clean:
	@echo "Cleaning up..."
	@$(GO_CLEAN)
	@rm -rf $(BIN_DIR)
	@echo "Cleanup complete!"

help:
	@echo "Makefile Commands:"
	@echo "  make build        - Build the CLI"
	@echo "  make build-windows - Build Windows binary"
	@echo "  make build-linux   - Build Linux binary"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean up build files"
