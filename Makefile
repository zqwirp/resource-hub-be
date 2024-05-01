# Makefile for a Go project with multiple sub-directories

# Default GOOS and GOARCH
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# Binary name based on the operating system
ifeq ($(GOOS),windows)
	BINARY_NAME=reshub.exe
else
	BINARY_NAME=reshub
endif

MAIN_PATH=./cmd/server
BINARY_PATH=./bin/$(BINARY_NAME)

# Automatically find all go files for dependency checking
SOURCES := $(shell find . -name '*.go')

# Default target
all: build

# Phony targets for commands that do not create files
.PHONY: build run test clean lint

# Ensure the bin directory exists
$(shell mkdir -p bin)

# Build the binary
build: $(BINARY_PATH)

# Rule to build the main application binary
$(BINARY_PATH): $(SOURCES)
	@echo "Building for $(GOOS)/$(GOARCH)..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY_PATH) $(MAIN_PATH)

# Run the application
run: build
	@echo "Running..."
	./$(BINARY_PATH)

# Test all go packages
# test:
# 	@echo "Testing..."
# 	go test ./...

# Clean up binaries and the bin directory
clean:
	@echo "Cleaning up..."
	go clean
	rm -rf ./bin

# Install dependencies and tidy up
deps:
	@echo "Installing dependencies..."
	go mod tidy

# Lint the project
# lint:
# 	@echo "Linting the code..."
# 	golint ./...
