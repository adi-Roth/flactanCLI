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

# Target architectures
ARCHS=amd64 arm64

# Build all targets
all: build-linux build-windows build-mac

lint:
	@echo "Running linter..."
	@golangci-lint run ./... -v
	@echo "Linter complete!"

build:
	@echo "Building $(PROJECT_NAME)..."
	@mkdir -p $(BIN_DIR)
	@$(GO_BUILD) -o $(BIN_DIR)/$(PROJECT_NAME) $(SRC_DIR)/main.go
	@echo "Build complete!"

build-windows:
	@echo "🚀 Building Windows binaries..."
	@mkdir -p $(BIN_DIR)
	GOOS=windows go get -d ./...  # Ensure Windows dependencies are downloaded
	@for arch in $(ARCHS); do \
		GOOS=windows GOARCH=$$arch $(GO_BUILD) -o $(BIN_DIR)/$(PROJECT_NAME)-windows-$$arch.exe $(SRC_DIR)/main.go; \
		echo "@echo off\nmove /Y $(PROJECT_NAME)-windows-$$arch.exe \"C:\\Program Files\\$(PROJECT_NAME).exe\"\necho ✅ Installed $(PROJECT_NAME) to C:\\Program Files\\" > $(BIN_DIR)/install-windows-$$arch.bat; \
		zip -j $(BIN_DIR)/$(PROJECT_NAME)-windows-$$arch.zip $(BIN_DIR)/$(PROJECT_NAME)-windows-$$arch.exe $(BIN_DIR)/install-windows-$$arch.bat; \
		echo "✅ Packaged: $(BIN_DIR)/$(PROJECT_NAME)-windows-$$arch.zip"; \
	done

build-linux:
	@echo "🚀 Building Linux binaries..."
	@mkdir -p $(BIN_DIR)
	@for arch in $(ARCHS); do \
		GOOS=linux GOARCH=$$arch $(GO_BUILD) -o $(BIN_DIR)/$(PROJECT_NAME)-linux-$$arch $(SRC_DIR)/main.go; \
		echo "#!/bin/bash\nsudo mv $(PROJECT_NAME)-linux-$$arch /usr/local/bin/$(PROJECT_NAME)\necho '✅ Installed $(PROJECT_NAME) to /usr/local/bin/'" > $(BIN_DIR)/install-linux-$$arch.sh; \
		chmod +x $(BIN_DIR)/install-linux-$$arch.sh; \
		zip -j $(BIN_DIR)/$(PROJECT_NAME)-linux-$$arch.zip $(BIN_DIR)/$(PROJECT_NAME)-linux-$$arch $(BIN_DIR)/install-linux-$$arch.sh; \
		echo "✅ Packaged: $(BIN_DIR)/$(PROJECT_NAME)-linux-$$arch.zip"; \
	done

build-mac:
	@echo "🚀 Building macOS binaries..."
	@mkdir -p $(BIN_DIR)
	@for arch in $(ARCHS); do \
		GOOS=darwin GOARCH=$$arch $(GO_BUILD) -o $(BIN_DIR)/$(PROJECT_NAME)-mac-$$arch $(SRC_DIR)/main.go; \
		echo "#!/bin/bash\nsudo mv $(PROJECT_NAME)-mac-$$arch /usr/local/bin/$(PROJECT_NAME)\necho '✅ Installed $(PROJECT_NAME) to /usr/local/bin/'" > $(BIN_DIR)/install-mac-$$arch.sh; \
		chmod +x $(BIN_DIR)/install-mac-$$arch.sh; \
		zip -j $(BIN_DIR)/$(PROJECT_NAME)-mac-$$arch.zip $(BIN_DIR)/$(PROJECT_NAME)-mac-$$arch $(BIN_DIR)/install-mac-$$arch.sh; \
		echo "✅ Packaged: $(BIN_DIR)/$(PROJECT_NAME)-mac-$$arch.zip"; \
	done

setup-cli:

test:
	@echo "Running tests..."
	@$(GO_TEST) -v ./$(TEST_DIR)/... --coverpkg=./cmd,./internal/... -coverprofile=coverage.out
	@echo "All tests passed!"

show-coverage:
	@echo "Showing coverage..."
	@$(GO) tool cover -html=coverage.out

clean:
	@echo "Cleaning up..."
	@$(GO_CLEAN)
	@rm -rf $(BIN_DIR)
	@rm -f coverage.out
	@echo "Cleanup complete!"

help:
	@echo "Makefile Commands:"
	@echo "  make build        - Build the CLI"
	@echo "  make build-windows - Build Windows binary"
	@echo "  make build-linux   - Build Linux binary"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean up build files"
